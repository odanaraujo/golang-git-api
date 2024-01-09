package converter

import (
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/model/repository/entity"
)

func ConverterEntitytoDomain(userEntity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUSerDomain(userEntity.Name, userEntity.Email, userEntity.Password, userEntity.Age)
	domain.SetID(userEntity.ID.Hex())
	return domain
}
