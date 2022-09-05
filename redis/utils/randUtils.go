package utils

import (
	"math/rand"
	"strings"
)

const CHARS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// 随机生成字符串 ([a-zA-Z0-9])
func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(CHARS)
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteByte(l)
	}
	return str.String()
}

// 随机生成数字字符串 ([0-9])
func RandNumString(lenNum int) string {
	str := strings.Builder{}
	length := 10
	for i := 0; i < lenNum; i++ {
		str.WriteByte(CHARS[52+rand.Intn(length)])
	}
	return str.String()
}

// 随机生成字母字符串 ([a-zA-Z])
func RandString(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteByte(CHARS[rand.Intn(length)])
	}
	return str.String()
}
