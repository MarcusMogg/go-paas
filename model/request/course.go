package request

// CourseIDReq 课程ID
type CourseIDReq struct {
	ID uint `json:"cid" form:"cid" binding:"required"`
}

// CourseReq 创建课程申请
type CourseReq struct {
	//TeacherID   uint   `json:"id" binding:"required"`
	Instruction string `json:"instruction" binding:"required"`
	Name        string `json:"name" binding:"required"`
}

// CourseUReq 课程更新申请
type CourseUReq struct {
	CourseIDReq
	CourseReq
}
