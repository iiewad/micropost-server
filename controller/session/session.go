package session

import (
	"log"

	"github.com/gin-gonic/gin"
)

type sessionNewParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

func New(c *gin.Context) {
	var sessionNewParams sessionNewParams
	c.BindJSON(&sessionNewParams)
	log.Println(sessionNewParams)
}
