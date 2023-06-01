package rosm

import (
	"bytes"
	"strconv"
)

// StrIsTrue 判断中文的真假
func StrIsTrue(str string) bool {
	return str == "True" || str == "true"
}

// Ftoone 保留一位小数并转化string
func Ftoone(f float64) string {
	if f == 0 {
		return "0"
	}
	return strconv.FormatFloat(f, 'f', 1, 64)
}

// 在字符串插入字符串
func InstStringN(s, substr string, n int) string {
	if len(s) == 0 {
		return s
	}
	var buffer bytes.Buffer
	for i, c := range []rune(s) {
		if i > 0 && i%n == 0 {
			buffer.WriteString(substr)
		}
		buffer.WriteRune(c)
	}
	return buffer.String()
}
