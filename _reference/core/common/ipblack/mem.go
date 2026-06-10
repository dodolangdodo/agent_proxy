package ipblack

import (
	"sync"
	"time"
)

var ipBlackMap = &sync.Map{}

func memSetIPBlack(ip string, duration time.Duration) {
	newExpiry := time.Now().Add(duration)
	newExpiryPtr := &newExpiry

	v, loaded := ipBlackMap.LoadOrStore(ip, newExpiryPtr)
	if loaded {
		expiredAtPtr, ok := v.(*time.Time)
		if !ok {
			// Type assertion failed, replace with new value
			ipBlackMap.Store(ip, newExpiryPtr)
			return
		}

		// If current value is expired, replace it with new value
		if time.Now().After(*expiredAtPtr) {
			ipBlackMap.CompareAndSwap(ip, expiredAtPtr, newExpiryPtr)
		}
	}
}

func memGetIPIsBlock(ip string) bool {
	v, ok := ipBlackMap.Load(ip)
	if !ok {
		return false
	}

	expiredAtPtr, ok := v.(*time.Time)
	if !ok {
		return false
	}

	if time.Now().After(*expiredAtPtr) {
		ipBlackMap.CompareAndDelete(ip, expiredAtPtr)
		return false
	}

	return true
}
