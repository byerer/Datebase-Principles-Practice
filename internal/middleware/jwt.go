package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var jwtKey = []byte("your_secret_key")

type Payload struct {
	UserId uint `json:"user_id"`
}

type Claims struct {
	userID int64
	jwt.RegisteredClaims
}

func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录"})
			c.Abort()
			return
		}
		claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录"})
			c.Abort()
			return
		}
		c.Set("userID", claims.userID)
		c.Next()
	}
}

// GenerateJWT 生成 JWT 字符串
func GenerateJWT(userID int64) (string, error) {
	claims := Claims{
		userID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "GradingSystem",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

// ParseToken 解析 JWT 字符串并验证
func ParseToken(tokenString string) (*Claims, error) {
	// 解析 token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// 验证 token 并返回结果
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
