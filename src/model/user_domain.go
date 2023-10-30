package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
)

// constructor interface
func NewUSerDomain(name, email, password string, age uint8) UserDomainInterface {
	return &UserDomain{
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
	}
}

type UserDomain struct {
	Name     string
	Email    string
	Password string
	Age      uint8
}

func (domain *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(domain.Password))
	domain.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *exception.Exception
	UpdateUser(string) *exception.Exception
	FindUser(string) (*UserDomain, *exception.Exception)
	DeleteUser(string) *exception.Exception
}
