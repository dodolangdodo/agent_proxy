package common

import (
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/labring/aiproxy/core/common/conv"
)

func ShortUUID() string {
	var buf [32]byte

	bytes := uuid.New()
	hex.Encode(buf[:], bytes[:])

	return conv.BytesToString(buf[:])
}
