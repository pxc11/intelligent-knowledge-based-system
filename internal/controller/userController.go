package controller

import (
	"ikbs/internal/model"
	"ikbs/lib/JWT"
	"ikbs/lib/db"
	"ikbs/lib/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Username string `json:"username" label:"用户名" binding:"required"`
	Password string `json:"password" label:"密码" binding:"required"`
}

func Login(c *gin.Context) {
	req, isSuccess := ValidateRequest[LoginReq](c)
	if !isSuccess {
		return
	}

	var user model.User
	if err := db.GetDb().Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"msg": "用户名或密码错误"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(400, gin.H{"msg": "用户名或密码错误"})
		return
	}

	err = JWT.GenerateTokenCookie(c, int64(user.ID))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(500, gin.H{"msg": "系统错误"})
	} else {
		c.JSON(200, gin.H{"msg": "登录成功"})
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

type RegisterReq struct {
	Username string `json:"username" label:"用户名" binding:"required,alphanum,min=5,max=20"`
	Password string `json:"password" label:"密码" binding:"required,min=6,max=64"`
}

func Register(c *gin.Context) {
	req, isSuccess := ValidateRequest[RegisterReq](c)
	if !isSuccess {
		return
	}
	var count int64
	if err := db.GetDb().Model(&model.User{}).
		Where("username = ?", req.Username).
		Count(&count).Error; err != nil {
		logger.Error(err.Error())
		c.JSON(500, gin.H{"msg": "系统错误"})
		return
	}

	if count > 0 {
		c.JSON(400, gin.H{"msg": "用户名已经存在"})
		return
	}

	// 2. 密码加密（必须）
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(500, gin.H{"msg": "系统错误1"})
		return
	}

	// 3. 创建用户
	user := model.User{
		Username: req.Username,
		Password: string(hashed),
	}

	if err := db.GetDb().Create(&user).Error; err != nil {
		logger.Error(err.Error())
		c.JSON(500, gin.H{"msg": "注册失败"})
		return
	}

	// 4. 返回结果（不要返回 password）
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})

}
