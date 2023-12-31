package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-course/api/models"
)

func Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Save()
	c.JSON(
		http.StatusCreated,
		gin.H{"message": "User created successfully!"},
	)
}
