package api

import (
	"fmt"
	"paas/middleware"
	"paas/model/entity"
	"paas/model/request"
	"paas/model/response"
	"paas/service"
	"paas/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var r request.RegisterData
	if err := c.BindJSON(&r); err == nil && utils.MatchEmail(r.Email) && utils.MatchStudentID(r.UserName) {
		savePath := fmt.Sprintf("%d.png", time.Now().Unix())
		utils.DrawText([]rune(r.UserName)[0], "source/avator/0/", savePath)
		user := &entity.MUser{
			UserName: r.UserName,
			NickName: r.NickName,
			Email:    r.Email,
			Password: r.Password,
			Role:     entity.Student,
			Avatar:   "source/avator/0/" + savePath,
		}

		if err = service.Register(user); err == nil {
			response.OkWithMessage("注册成功", c)
		} else {
			response.FailWithMessage(fmt.Sprintf("%v", err), c)
		}

	} else {
		response.FailValidate(c)
	}
}

// Login 用户登录
func Login(c *gin.Context) {
	var r request.LoginData
	if err := c.BindJSON(&r); err == nil {
		user := &entity.MUser{UserName: r.UserName, Password: r.Password}

		if service.Login(user) {
			tokenNext(c, user)
		} else {
			response.FailWithMessage("账号或者密码错误", c)
		}

	} else {
		response.FailValidate(c)
	}
}

func tokenNext(c *gin.Context, u *entity.MUser) {
	j := middleware.NewJWT()
	claim := middleware.JWTClaim{
		UserID:   u.ID,
		UserName: u.UserName,
		Role:     u.Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 60*60*24*7,
			Issuer:    "715worker",
		},
	}
	token, err := j.CreateToken(claim)
	if err != nil {
		response.FailWithMessage("token创建失败", c)
		return
	}
	response.OkWithData(token, c)
}

// GetUserInfoByID 获取指定用户信息
func GetUserInfoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	u, err := service.GetUserInfoByID(uint(id))
	if err == nil {
		response.OkWithData(u, c)
	} else {
		response.Fail(c)
	}
}

// UpdateEmail 获取用户邮箱
func UpdateEmail(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var ur request.EmailReq
	if err := c.BindJSON(&ur); err == nil {
		user.Email = ur.Email
		if err = service.UpdateUser(user); err == nil {
			response.Ok(c)
		} else {
			response.FailWithMessage(err.Error(), c)
		}
	} else {
		response.FailValidate(c)
	}
}

// UpdateAvatar 上传头像
func UpdateAvatar(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	savePath := "source/avator/" + fmt.Sprintf("%d/", user.ID)
	fileName, suf, err := uploadFile(savePath, c)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}
	user.Avatar = fmt.Sprintf("%s%s%s", savePath, fileName, suf)
	if err = service.UpdateUser(user); err == nil {
		response.Ok(c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// UpdatePassword 获取用户邮箱
func UpdatePassword(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var ur request.PasswordReq
	if err := c.BindJSON(&ur); err == nil {
		user.Password = utils.AesEncrypt(ur.Password)
		if err = service.UpdateUser(user); err == nil {
			response.Ok(c)
		} else {
			response.FailWithMessage(err.Error(), c)
		}
	} else {
		response.FailValidate(c)
	}
}
