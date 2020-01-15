package controllers

import (
	"github.com/gin-gonic/gin"
	base "github.com/grestful/core"
	"{{package}}/models"
	"{{package}}/services"
)

func UserIndex(c *gin.Context) {
	base.RunProcess(&services.UserIdService{}, c)
}

func UserList(c *gin.Context) {
	p:=&services.UserIdService{}
	print(p)
	base.RunProcess(p, c)
	//base.RunProcess(&services.UserService{}, c)
}

func SaveUser(c *gin.Context) {
	p := &services.UserSaveService{}
	p.Req = &services.UserSaveRequest{
		DefaultModel: &models.DefaultModel{}, Type: "",
	}
	p.ProcessFun = func(u *base.Controller) base.IError {
		u.Data = "hello world"
		return nil
	}
	base.RunProcess(p, c)
}
