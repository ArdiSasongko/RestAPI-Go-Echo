package web

type UserReq struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserUpdateReq struct {
	Name     string `json:"name"`
	Email    string `validate:"email" json:"email"`
	Password string `json:"password"`
}

type UserLoginReq struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
