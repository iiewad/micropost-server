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

var result models.Result

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
