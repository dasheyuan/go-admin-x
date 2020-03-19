package util

import (
	_ "go-admin-x/internal/util/conf" // init conf
	"go-admin-x/internal/util/log"
)

// GatherMetrics 收集一些被动指标
func GatherMetrics() {
	//mc.GatherMetrics()
	//redis.GatherMetrics()
	//db.GatherMetrics()
}

// Reset all utils
func Reset() {
	log.Reset()
	//orm.Reset()
	//db.Reset()

	//mc.Reset()
}

// Stop all utils
func Stop() {
}
