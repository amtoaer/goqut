package utils

// CheckError 用于检测错误，出现错误则panic
func CheckError(err error, message string) {
	if err != nil {
		panic(message)
	}
}
