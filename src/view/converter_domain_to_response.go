package view

import (
	"github.com/odanaraujo/golang/users-api/src/controller/model/response"
	"github.com/odanaraujo/golang/users-api/src/model"
)

func ConverterDomainToResponse(user model.UserDomainInterface) *response.UserResponse {
	return &response.UserResponse{
		Name:  user.GetName(),
		Email: user.GetEmail(),
		Age:   user.GetAge(),
	}
}
