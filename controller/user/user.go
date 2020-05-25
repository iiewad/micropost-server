package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iiewad/micropost-server/common"
	"github.com/iiewad/micropost-server/models"
)

type registerParams struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,max=128,min=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required,max=128,min=6"`
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
	c.JSON(200, gin.H{"code": 0, "msg": "success", "data": userInfo})
}
