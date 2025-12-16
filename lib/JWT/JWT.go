package JWT

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"ikbs/lib/config"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// jwt token 生成
func GenerateToken(userId int64) (string, error) {
	appConfig := config.LoadConfig()

	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(appConfig.JWT.Expire * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(appConfig.JWT.Secret)
}

// jwt中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, err := c.Cookie("jwt-token")
		if err != nil {

		}
		if auth == "" {
			c.AbortWithStatusJSON(401, gin.H{"msg": "missing token"})
			return
		}

		// Bearer token
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"msg": "invalid token"})
			return
		}

		// 保存用户信息
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
