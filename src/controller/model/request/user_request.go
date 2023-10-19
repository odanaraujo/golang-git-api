package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=80"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%^&*()_+-="`
	Age      uint8  `json:"age" binding:"required,max=110"`
}
