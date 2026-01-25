package controller

import (
	"errors"
	"ikbs/internal/myValidator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 通用验证
func ValidateRequest[reqType any](c *gin.Context) (reqType, bool) {
	var req reqType
	err := c.ShouldBindJSON(&req)
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		c.JSON(200, gin.H{
			"msg": errs[0].Translate(myValidator.GetTrans()),
			"sts": false,
		})
		return req, false
	}
	if err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
			"sts": false},
		)
		return req, false
	}
	return req, true
}
