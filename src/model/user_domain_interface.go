package model

type UserDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() uint8
	SetID(id string)

	EncryptPassword()
}

// constructor interface
func NewUSerDomain(name, email, password string, age uint8) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
		age:      age,
	}
}

func NewUSerUpdateDomain(name string, age uint8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(email string, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
