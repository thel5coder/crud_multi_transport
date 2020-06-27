package request

type UserRequest struct {
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"min=6"`
	MobilePhone string `json:"mobile_phone" validate:"required"`
}
