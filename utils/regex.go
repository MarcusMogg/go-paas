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

// MatchStudentID 判断学生ID,必须以数字字母开头，只能包含数字字母和下划线
func MatchStudentID(str string) bool {
	return match(`^[A-Za-z0-9][a-zA-Z0-9_]*$`, str)
}
