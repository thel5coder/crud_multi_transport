package viewmodel

type UserVm struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobile_phone"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
