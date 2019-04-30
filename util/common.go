package util

import (
	"bytes"
)

// CalPage 计算总页数.
func CalPage(total uint64, pageSize int) int {
	// 计算总页数
	if m := total % uint64(pageSize); m == 0 {
		return int(total / uint64(pageSize))
	}
	return int(total/uint64(pageSize) + 1)
}

// BufferWriteString method used to splice newline string
func BufferWriteString(args ...string) string {
	var buf bytes.Buffer
	for index, arg := range args {
		buf.WriteString(arg)
		if index+1 != len(args) {
			buf.WriteString("<br/>")
		}
	}
	return buf.String()
}
