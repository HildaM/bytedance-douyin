package utils

import "time"

func NowTimestamp() int64 {
	return time.Now().UnixMilli()
}

func NowDurationTimestamp(duration time.Duration) int64 {
	return time.Now().Add(duration).UnixMilli()
}
