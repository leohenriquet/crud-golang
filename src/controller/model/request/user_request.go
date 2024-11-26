package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%"`
	Age      int8   `json:"age" binding:"required,min=0,max=140"`
}
