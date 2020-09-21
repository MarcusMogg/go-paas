package request

// AddStudentsReq 教师添加学生
type AddStudentsReq struct {
	CourseIDReq
	UserNames []string `json:"ids" binding:"required"`
}

// DelStudentReq 教师同意/拒绝学生加入
type DelStudentReq struct {
	CourseIDReq
	GetByID
}
