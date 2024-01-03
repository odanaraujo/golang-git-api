package converter

import (
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/model/repository/entity"
)

func ConverterDomainToEntity(userDomain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Name:     userDomain.GetName(),
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Age:      userDomain.GetAge(),
	}
}
