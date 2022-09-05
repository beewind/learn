package utils

// 是否无效
func IsPhoneInvalid(phone string) bool {
	return false
}
func IsNumInvalid(num string) bool {
	for _, v := range num {
		if v < '0' || v > '9' {
			return true
		}
	}
	return false
}
