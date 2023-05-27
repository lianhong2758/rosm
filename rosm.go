package rosm

import (
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
