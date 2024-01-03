package model

type UserDomainInterface interface {
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
