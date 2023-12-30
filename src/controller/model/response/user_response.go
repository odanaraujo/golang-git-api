package response

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}
