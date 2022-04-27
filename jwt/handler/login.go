package handler

import (
	"codebase/jwt/model"

	"github.com/gin-gonic/gin"
)

var User = map[string]string{
	"test":  "hello1",
	"admin": "admin",
}

type LoginResp struct {
	ID    string
	Token string
}

func Login(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	if id == "" || password == "" {
		ResponseFail(c, 400, "phone_num not exit")
		return
	}
	if pw, ok := User[id]; !ok || pw != password {
		ResponseFail(c, 401, "账号或密码错误")
		return
	}
	token := model.GenerateToken(id)
	ResponseSuccess(c, &LoginResp{ID: id, Token: token})
}
