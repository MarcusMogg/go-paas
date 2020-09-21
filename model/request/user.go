package request

// RegisterData 注册时传入参数
type RegisterData struct {
	UserName string `form:"username" json:"username" binding:"required"`
	NickName string `form:"nickname" json:"nickname" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// LoginData 登录时传入参数
type LoginData struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// EmailReq 修改用户邮箱时传入参数
type EmailReq struct {
	Email string `json:"email" binding:"required"`
}

// PasswordReq 修改用户密码时传入参数
type PasswordReq struct {
	Password string `json:"password" binding:"required"`
}
