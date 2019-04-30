package util

import (
	"crypto/sha1"
	"fmt"
)

// Sha1 生成sha1加密串.
func Sha1(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	h := sha1.New()
	h.Write(b)
	bs := h.Sum(nil)
	disc := fmt.Sprintf("%x", bs)
	return disc
}
