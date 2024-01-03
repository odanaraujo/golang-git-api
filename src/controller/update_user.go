package controller

import "github.com/gin-gonic/gin"

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "updateUser",
	})
}
