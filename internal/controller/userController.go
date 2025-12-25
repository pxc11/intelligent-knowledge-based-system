package controller

import (
	"ikbs/internal/model"
	"ikbs/internal/myValidator"
	"ikbs/lib/JWT"
	"ikbs/lib/db"
	"ikbs/lib/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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

type RegisterReq struct {
	Username string `json:"username" binding:"required,alphanum,min=5,max=20"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}

func Register(c *gin.Context) {
	var req RegisterReq
	err := c.ShouldBindJSON(&req)
	if errs, ok := err.(validator.ValidationErrors); ok {
		c.JSON(200, gin.H{
			"msg": errs[0].Translate(myValidator.GetTrans()),
			"sts": false,
		})
		return
	}

	if err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
			"sts": false},
		)
		return
	}

	var count int64
	if err := db.GetDb().Model(&model.User{}).
		Where("username = ?", req.Username).
		Count(&count).Error; err != nil {
		logger.Error(err.Error())
		c.JSON(200, gin.H{"sts": false, "msg": "系统错误"})
		return
	}

	if count > 0 {
		c.JSON(200, gin.H{"sts": false, "msg": "用户名已经存在"})
		return
	}

	// 2. 密码加密（必须）
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(200, gin.H{"msg": "系统错误1"})
		return
	}

	// 3. 创建用户
	user := model.User{
		Username: req.Username,
		Password: string(hashed),
	}

	if err := db.GetDb().Create(&user).Error; err != nil {
		logger.Error(err.Error())
		c.JSON(200, gin.H{"sts": false, "msg": "注册失败"})
		return
	}

	// 4. 返回结果（不要返回 password）
	c.JSON(200, gin.H{
		"sts": false,
		"msg": "注册成功",
	})

}
