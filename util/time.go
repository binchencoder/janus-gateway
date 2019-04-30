package util

import (
	"time"
)

// MsStamp 返回毫秒时间戳.
func MsStamp() int64 {
	return time.Now().UnixNano() / 1000000
}
