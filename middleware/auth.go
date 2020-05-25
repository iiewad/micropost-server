package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iiewad/micropost-server/config"
	"github.com/iiewad/micropost-server/models"
)

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

// Auth Middleware
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := models.Result{
			Code: http.StatusUnauthorized,
			Msg:  "无法认证, 重新登录",
			Data: nil,
		}
		auth := context.Request.Header.Get("Authorization")
		if len(auth) != 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{"result": result})
		}
		auth = strings.Fields(auth)[1]
		_, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result.Msg = "token 过期" + err.Error()
			context.JSON(http.StatusUnauthorized, gin.H{"result": result})
		} else {
			log.Println("Token 正确")
		}
		context.Next()
	}
}
