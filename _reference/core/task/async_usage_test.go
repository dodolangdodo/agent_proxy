//nolint:testpackage
package task

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/labring/aiproxy/core/common/balance"
	"github.com/labring/aiproxy/core/model"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type failingAsyncUsageBalance struct {
	err error
}

func (b failingAsyncUsageBalance) GetGroupRemainBalance(
	context.Context,
	model.GroupCache,
) (float64, balance.PostGroupConsumer, error) {
	return 100, failingAsyncUsageConsumer(b), nil
}

func (b failingAsyncUsageBalance) GetGroupQuota(
	context.Context,
	model.GroupCache,
) (*balance.GroupQuota, error) {
	return &balance.GroupQuota{Total: 100, Remain: 100}, nil
}

type failingAsyncUsageConsumer struct {
	err error
}

func (c failingAsyncUsageConsumer) PostGroupConsume(
	context.Context,
	string,
	float64,
) (float64, error) {
	return 0, c.err
}

type preChargeFailingAsyncUsageBalance struct {
	err error
}

func (b preChargeFailingAsyncUsageBalance) GetGroupRemainBalance(
	context.Context,
	model.GroupCache,
) (float64, balance.PostGroupConsumer, error) {
	return 0, nil, b.err
}

func (b preChargeFailingAsyncUsageBalance) GetGroupQuota(
	context.Context,
	model.GroupCache,
) (*balance.GroupQuota, error) {
	return nil, b.err
}

func TestCompleteAsyncUsageIgnoresMissingLog(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Log{}, &model.AsyncUsageInfo{}))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	info := &model.AsyncUsageInfo{
		RequestID:       "missing_log",
		RequestAt:       time.Now(),
		Status:          model.AsyncUsageStatusPending,
		Model:           "gpt-5.4",
		ProcessingToken: "claim-token",
	}
	require.NoError(t, model.CreateAsyncUsageInfo(info))

	usage := model.Usage{
		InputTokens: 10,
		TotalTokens: 10,
	}

	require.NoError(t, completeAsyncUsage(context.Background(), info, usage, model.UsageContext{}))
	require.Equal(t, model.AsyncUsageStatusCompleted, info.Status)
	require.Equal(t, usage.InputTokens, info.Usage.InputTokens)
}

func TestCompleteAsyncUsageReturnsLogUpdateError(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Log{}, &model.AsyncUsageInfo{}))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	sqlDB, err := db.DB()
	require.NoError(t, err)
	require.NoError(t, sqlDB.Close())

	info := &model.AsyncUsageInfo{
		RequestID: "log_update_error",
		RequestAt: time.Now(),
		Status:    model.AsyncUsageStatusPending,
		Model:     "gpt-5.4",
	}
	usage := model.Usage{
		InputTokens: 10,
		TotalTokens: 10,
	}

	require.Error(t, completeAsyncUsage(context.Background(), info, usage, model.UsageContext{}))
	require.Equal(t, model.AsyncUsageStatusPending, info.Status)
	require.Equal(t, model.ZeroNullInt64(0), info.Usage.InputTokens)
}

func TestCompleteAsyncUsageRecordsBalanceConsumeErrorWithoutRetry(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(
		&model.Log{},
		&model.AsyncUsageInfo{},
		&model.ConsumeError{},
	))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	oldBalance := balance.Default
	balance.Default = failingAsyncUsageBalance{err: errors.New("balance unavailable")}
	t.Cleanup(func() {
		balance.Default = oldBalance
	})

	require.NoError(t, model.CacheSetGroup(&model.GroupCache{
		ID:     "group-async-balance",
		Status: model.GroupStatusEnabled,
	}))
	t.Cleanup(func() {
		require.NoError(t, model.CacheDeleteGroup("group-async-balance"))
	})

	info := &model.AsyncUsageInfo{
		RequestID:  "balance_error",
		RequestAt:  time.Now(),
		Status:     model.AsyncUsageStatusPending,
		Model:      "gpt-5.4",
		GroupID:    "group-async-balance",
		TokenID:    1,
		TokenName:  "token-1",
		Price:      model.Price{InputPrice: 1, InputPriceUnit: 1},
		UpstreamID: "resp_balance_error",
	}
	require.NoError(t, model.CreateAsyncUsageInfo(info))

	err = completeAsyncUsage(context.Background(), info, model.Usage{
		InputTokens: 10,
		TotalTokens: 10,
	}, model.UsageContext{})
	require.NoError(t, err)
	require.Equal(t, model.AsyncUsageStatusCompleted, info.Status)
	require.False(t, info.BalanceConsumed)

	var consumeErrors []model.ConsumeError
	require.NoError(t, db.Find(&consumeErrors).Error)
	require.Len(t, consumeErrors, 1)
	require.Equal(t, "balance_error", consumeErrors[0].RequestID)
	require.Equal(t, float64(10), consumeErrors[0].UsedAmount)

	var got model.AsyncUsageInfo
	require.NoError(t, db.First(&got, info.ID).Error)
	require.Equal(t, model.AsyncUsageStatusCompleted, got.Status)
	require.False(t, got.BalanceConsumed)
	require.Equal(t, float64(10), got.Amount.UsedAmount)
}

func TestCompleteAsyncUsageRetriesPreChargeBalanceError(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(
		&model.Log{},
		&model.AsyncUsageInfo{},
		&model.ConsumeError{},
	))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	oldBalance := balance.Default
	balance.Default = preChargeFailingAsyncUsageBalance{
		err: errors.New("balance lookup unavailable"),
	}
	t.Cleanup(func() {
		balance.Default = oldBalance
	})

	require.NoError(t, model.CacheSetGroup(&model.GroupCache{
		ID:     "group-pre-charge-balance",
		Status: model.GroupStatusEnabled,
	}))
	t.Cleanup(func() {
		require.NoError(t, model.CacheDeleteGroup("group-pre-charge-balance"))
	})

	info := &model.AsyncUsageInfo{
		RequestID:  "pre_charge_balance_error",
		RequestAt:  time.Now(),
		Status:     model.AsyncUsageStatusPending,
		Model:      "gpt-5.4",
		GroupID:    "group-pre-charge-balance",
		TokenID:    1,
		TokenName:  "token-1",
		Price:      model.Price{InputPrice: 1, InputPriceUnit: 1},
		UpstreamID: "resp_pre_charge_balance_error",
	}
	require.NoError(t, model.CreateAsyncUsageInfo(info))

	err = completeAsyncUsage(context.Background(), info, model.Usage{
		InputTokens: 10,
		TotalTokens: 10,
	}, model.UsageContext{})
	require.ErrorContains(t, err, "consume async usage balance before charge")
	require.Equal(t, model.AsyncUsageStatusPending, info.Status)
	require.False(t, info.BalanceConsumed)

	var consumeErrors []model.ConsumeError
	require.NoError(t, db.Find(&consumeErrors).Error)
	require.Empty(t, consumeErrors)

	var got model.AsyncUsageInfo
	require.NoError(t, db.First(&got, info.ID).Error)
	require.Equal(t, model.AsyncUsageStatusPending, got.Status)
	require.False(t, got.BalanceConsumed)
	require.Equal(t, float64(0), got.Amount.UsedAmount)
}

func TestCompleteAsyncUsagePreservesStoredPriceCondition(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Log{}, &model.AsyncUsageInfo{}))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	requestID := "async_condition"
	require.NoError(t, db.Create(&model.Log{
		RequestID:        model.EmptyNullString(requestID),
		AsyncUsageStatus: model.AsyncUsageStatusPending,
	}).Error)

	info := &model.AsyncUsageInfo{
		RequestID: requestID,
		RequestAt: time.Now(),
		Status:    model.AsyncUsageStatusPending,
		Model:     "video-model",
		Price: model.Price{
			OutputPrice:     0.1,
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						OutputPrice:     0.4,
						OutputPriceUnit: 1,
					},
				},
			},
		},
		UsageContext: model.UsageContext{
			Resolution: "720P",
		},
		ProcessingToken: "claim-token",
	}
	require.NoError(t, model.CreateAsyncUsageInfo(info))

	require.NoError(t, completeAsyncUsage(context.Background(), info, model.Usage{
		OutputTokens: 5,
		TotalTokens:  5,
	}, model.UsageContext{}))
	require.Equal(t, model.AsyncUsageStatusCompleted, info.Status)
	require.Equal(t, "720P", info.UsageContext.Resolution)
	require.Equal(t, 2.0, info.Amount.UsedAmount)

	var got model.Log
	require.NoError(t, db.Where("request_id = ?", requestID).First(&got).Error)
	require.Equal(t, "720P", got.UsageContext.Resolution)
	require.Equal(t, model.ZeroNullFloat64(0.4), got.Price.OutputPrice)
	require.Equal(t, model.ZeroNullInt64(1), got.Price.OutputPriceUnit)
	require.Empty(t, got.Price.ConditionalPrices)
	require.Equal(t, 2.0, got.Amount.UsedAmount)
}

func TestCompleteAsyncUsageChargesStoredPerRequestPrice(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Log{}, &model.AsyncUsageInfo{}))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	requestID := "async_per_request"
	require.NoError(t, db.Create(&model.Log{
		RequestID:        model.EmptyNullString(requestID),
		AsyncUsageStatus: model.AsyncUsageStatusPending,
	}).Error)

	info := &model.AsyncUsageInfo{
		RequestID:       requestID,
		RequestAt:       time.Now(),
		Status:          model.AsyncUsageStatusPending,
		Model:           "video-model",
		Price:           model.Price{PerRequestPrice: 0.25, OutputPrice: 0.5, OutputPriceUnit: 1},
		ProcessingToken: "claim-token",
	}
	require.NoError(t, model.CreateAsyncUsageInfo(info))

	require.NoError(t, completeAsyncUsage(context.Background(), info, model.Usage{
		OutputTokens: 4,
		TotalTokens:  4,
	}, model.UsageContext{}))
	require.Equal(t, model.AsyncUsageStatusCompleted, info.Status)
	require.Equal(t, 0.25, info.Amount.UsedAmount)

	var got model.Log
	require.NoError(t, db.Where("request_id = ?", requestID).First(&got).Error)
	require.Equal(t, 0.25, got.Amount.UsedAmount)
}

func TestCompleteAsyncUsagePersistsBalanceConsumed(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Log{}, &model.AsyncUsageInfo{}))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	info := &model.AsyncUsageInfo{
		RequestID:       "async_balance_consumed",
		RequestAt:       time.Now(),
		Status:          model.AsyncUsageStatusPending,
		Model:           "video-model",
		BalanceConsumed: true,
		ProcessingToken: "claim-token",
	}
	require.NoError(t, model.CreateAsyncUsageInfo(info))

	require.NoError(t, completeAsyncUsage(context.Background(), info, model.Usage{
		OutputTokens: 4,
		TotalTokens:  4,
	}, model.UsageContext{}))
	require.Equal(t, model.AsyncUsageStatusCompleted, info.Status)
	require.True(t, info.BalanceConsumed)

	var got model.AsyncUsageInfo
	require.NoError(t, db.First(&got, info.ID).Error)
	require.True(t, got.BalanceConsumed)
	require.Equal(t, model.AsyncUsageStatusCompleted, got.Status)
}

func TestTouchAsyncUsagePollCursorAdvancesUpdatedAtAndNextPollAt(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.AsyncUsageInfo{}))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	oldUpdatedAt := time.Now().Add(-time.Hour)
	oldNextPollAt := time.Now().Add(-time.Minute)
	info := &model.AsyncUsageInfo{
		RequestID:       "pending",
		Status:          model.AsyncUsageStatusPending,
		UpdatedAt:       oldUpdatedAt,
		NextPollAt:      oldNextPollAt,
		Error:           "previous error",
		ProcessingToken: "claim-token",
	}
	require.NoError(t, db.Create(info).Error)

	beforeTouch := time.Now()

	touchAsyncUsagePollCursor(info)

	var got model.AsyncUsageInfo
	require.NoError(t, db.First(&got, info.ID).Error)
	require.True(t, got.UpdatedAt.After(oldUpdatedAt))
	require.True(t, got.NextPollAt.After(beforeTouch))
	require.Empty(t, got.Error)
	require.Empty(t, got.ProcessingToken)
}

func TestMarkAsyncUsageFailedWritesLogMessage(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Log{}, &model.AsyncUsageInfo{}))

	oldLogDB := model.LogDB
	model.LogDB = db
	t.Cleanup(func() {
		model.LogDB = oldLogDB
	})

	requestID := "async_fail_log"
	require.NoError(t, db.Create(&model.Log{
		RequestID:        model.EmptyNullString(requestID),
		AsyncUsageStatus: model.AsyncUsageStatusPending,
	}).Error)

	info := &model.AsyncUsageInfo{
		RequestID:       requestID,
		Status:          model.AsyncUsageStatusPending,
		ProcessingToken: "claim-token",
	}
	require.NoError(t, db.Create(info).Error)

	markAsyncUsageFailed(info, "upstream task failed")

	var got model.Log
	require.NoError(t, db.Where("request_id = ?", requestID).First(&got).Error)
	require.Equal(t, model.AsyncUsageStatusFailed, got.AsyncUsageStatus)
	require.Equal(t, "upstream task failed", string(got.Content))
}
