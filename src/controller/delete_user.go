package controller

import "github.com/gin-gonic/gin"

func DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "deleteUser",
	})
}
