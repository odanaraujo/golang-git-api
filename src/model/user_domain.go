package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() uint8

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

func (domain *userDomain) GetName() string {
	return domain.name
}

func (domain *userDomain) GetEmail() string {
	return domain.email
}

func (domain *userDomain) GetPassword() string {
	return domain.password
}
func (domain *userDomain) GetAge() uint8 {
	return domain.age
}

type userDomain struct {
	name     string
	email    string
	password string
	age      uint8
}

func (domain *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(domain.password))
	domain.password = hex.EncodeToString(hash.Sum(nil))
}
