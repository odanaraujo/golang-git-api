package controller

import "github.com/gin-gonic/gin"

func (uc *userControllerInterface) FindUserByEmail(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "getUserByEmail",
	})
}
