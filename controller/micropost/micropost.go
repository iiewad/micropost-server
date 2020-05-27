package micropost

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iiewad/micropost-server/common"
	"github.com/iiewad/micropost-server/middleware"
	"github.com/iiewad/micropost-server/models"
)

type micropostCP struct {
	Content string `json:"content"`
}
type micropostUP struct {
	UUID    string `json:"uuid"`
	Content string `json:"content"`
}

var result models.Result

// Update Micropost
func Update(c *gin.Context) {
	var micropostUP micropostUP
	if err := c.ShouldBindJSON(&micropostUP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var currentMicropost models.Micropost
	common.DB.Where("user_id = ? AND uuid = ?", middleware.CurrentUser.UUID, micropostUP.UUID).First(&currentMicropost)
	common.DB.Model(&currentMicropost).Update("content", &micropostUP.Content)
	result.Code = http.StatusOK
	result.Msg = "更新成功"
	result.Data = currentMicropost
	c.JSON(result.Code, gin.H{"result": result})
}

// List Micropost
func List(c *gin.Context) {
	var micropostList []models.Micropost
	common.DB.Where("user_id = ?", middleware.CurrentUser.UUID).Find(&micropostList)
	result.Code = http.StatusOK
	result.Msg = "success"
	result.Data = micropostList
	c.JSON(http.StatusOK, gin.H{"result": result})
}

// Create Micropost
func Create(c *gin.Context) {
	var micropostCP micropostCP
	if err := c.ShouldBindJSON(&micropostCP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newMicropost := models.Micropost{UserID: middleware.CurrentUser.UUID, Content: &micropostCP.Content}
	micropostID, err := models.AddMicropost(common.DB, newMicropost)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	common.DB.Where("uuid = ?", micropostID).First(&newMicropost)
	result.Code = http.StatusOK
	result.Msg = "创建成功"
	result.Data = newMicropost
	c.JSON(result.Code, gin.H{"result": result})
}
