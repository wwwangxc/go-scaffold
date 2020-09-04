package util

import (
	"crypto/sha256"
	"fmt"
)

// WithSecret 字符串加盐
func WithSecret(str string, salt string) string {
	s1 := sha256.New()
	s1.Write([]byte(str))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	return fmt.Sprintf("%x", s2.Sum(nil))
}
