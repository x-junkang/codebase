package handler

import (
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	ucid := c.GetString("id")
	if ucid == "" {
		ResponseFail(c, 10001, "id不存在")
		return
	}
	ResponseSuccess(c, ucid)
}
