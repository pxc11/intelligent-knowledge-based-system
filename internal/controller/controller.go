package controller

import (
	"ikbs/lib/JWT"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	err := JWT.GenerateTokenCookie(c, 1)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})

	} else {
		c.JSON(200, gin.H{"msg": "login success"})
	}
}

func GetUserInfo(c *gin.Context) {
	userInfo, exists := c.Get("user_info")
	if !exists {
		c.JSON(200, gin.H{"msg": "user_info not exist"})
	} else {
		c.JSON(200, gin.H{"msg": "get user info", "data": userInfo})
	}

}
