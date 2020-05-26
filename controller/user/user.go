package user

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iiewad/micropost-server/common"
	"github.com/iiewad/micropost-server/config"
	"github.com/iiewad/micropost-server/models"
	"golang.org/x/crypto/bcrypt"
)

type registerParams struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,max=128,min=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required,max=128,min=6"`
}

type loginParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=128,min=6"`
}

var result models.Result

// Login *
func Login(c *gin.Context) {
	var loginParams loginParams

	if err := c.ShouldBindJSON(&loginParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var user models.User
	common.DB.Where("email = ?", loginParams.Email).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginParams.Password))
	if err == nil {
		expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
		claims := jwt.StandardClaims{
			Audience:  user.Name,
			ExpiresAt: expiresTime,
			Id:        string(*user.UUID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "micropost-server",
			NotBefore: time.Now().Unix(),
			Subject:   "login",
		}
		var jwtSecret = []byte(config.Secret)
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err != nil {
			result.Code = http.StatusUnauthorized
			result.Msg = "Failed"
			result.Data = nil
			c.JSON(result.Code, gin.H{"result": result})
			return
		}
		result.Code = http.StatusOK
		result.Msg = "Success"
		result.Data = "Bearer " + token
		c.JSON(result.Code, gin.H{"result": result})
		return

	}
	result.Code = http.StatusBadRequest
	result.Msg = "登录失败"
	c.JSON(result.Code, gin.H{"result": result})
}

// Create User
func Create(c *gin.Context) {
	var registerParams registerParams
	if err := c.ShouldBindJSON(&registerParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if registerParams.Password != registerParams.PasswordConfirm {
		c.JSON(400, gin.H{"error": "密码不匹配"})
		return
	}
	user := models.User{Name: registerParams.Name, Email: registerParams.Email, PasswordHash: registerParams.Password}
	userID, err := models.AddUser(common.DB, user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var newUser models.User
	common.DB.Where("uuid = ?", userID).First(&newUser)
	var userInfo models.UserInfo
	userInfo = newUser.UserInfo()
	result.Code = http.StatusOK
	result.Msg = "创建成功"
	result.Data = userInfo
	c.JSON(result.Code, gin.H{"result": result})
}
