package utils

import "regexp"

func match(rep, str string) bool {
	reg, err := regexp.Compile(rep)
	if err != nil {
		return false // regexp error
	}
	return reg.MatchString(str)
}

// MatchEmail 判断是否是邮箱
func MatchEmail(str string) bool {
	return match(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, str)
}

// MatchStudentID 判断学生ID,必须为八位数字
func MatchStudentID(str string) bool {
	return match(`^[0-9]{8}$`, str)
}
