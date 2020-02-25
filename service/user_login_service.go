package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mithril/model"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Account  string `form:"account" json:"account" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) bool {
	var user model.User

	if err := model.DB.Where("account = ?", service.Account).First(&user).Error; err != nil {
		//return serializer.ParamErr("账号或密码错误", nil
		return false
	}

	if user.CheckPassword(service.Password) == false {
		//return serializer.ParamErr("账号或密码错误", nil)
		return false
	}

	// 设置session
	service.setSession(c, user)

	return true
}
