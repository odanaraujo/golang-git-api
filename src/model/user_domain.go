package model

type userDomain struct {
	id       string
	name     string
	email    string
	password string
	age      uint8
}

func (domain *userDomain) GetID() string {
	return domain.id
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

func (domain *userDomain) SetID(id string) {
	domain.id = id
}
