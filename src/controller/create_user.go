package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/controller/model/response"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/controller/model/request"
)

func CreateUser(ctx *gin.Context) {

	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		excp := exception.BadRequestException(fmt.Sprintf("there are some incorrect fields, error=%s", err))
		ctx.JSON(excp.Code, excp)
	}

	userResponse := response.UserResponse{
		ID:    strconv.FormatInt(int64(rand.Intn(10)+1), 10),
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	ctx.JSON(http.StatusOK, userResponse)
}
