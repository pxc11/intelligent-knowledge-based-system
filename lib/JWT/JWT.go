package JWT

import (
	"ikbs/lib/basic"
	"ikbs/lib/logger"
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
	jwtToken, err := token.SignedString([]byte(appConfig.JWT.Secret))
	if err != nil {
		return "", err
	}
	logger.Info("生成token", jwtToken)
	return jwtToken, nil
}

// jwt中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, err := c.Cookie("jwt-token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"msg": err.Error()})
			return
		}
		if auth == "" {
			c.AbortWithStatusJSON(401, gin.H{"msg": "missing token"})
			return
		}

		// Bearer token
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.LoadConfig().JWT.Secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"msg": err.Error()})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"msg": jwt.ErrTokenNotValidYet.Error()})
		}

		// 保存用户信息
		c.Set("user_info", claims)

		remaining := time.Until(claims.ExpiresAt.Time)
		if remaining < 10*time.Minute {

			err := GenerateTokenCookie(c, claims.UserId)
			if err != nil {
				logger.Error(err)
				return
			}

		}
		c.Next()

	}
}

func GenerateTokenCookie(c *gin.Context, userId int64) error {
	newToken, err := GenerateToken(userId)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	c.SetCookie("jwt-token", newToken, int(config.LoadConfig().JWT.Expire), "/", "", basic.IsSecure(c), true)
	return nil
}
