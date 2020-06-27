package transformator

type BrowseRequest struct {
	Search string `json:"search"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
}

type ReadRequest struct {
	ID string `json:"id"`
}


type UserRequest struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	MobilePhone string `json:"mobile_phone"`
}

type EditRequest struct {
	ID   string      `json:"id"`
	User UserRequest `json:"user"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}