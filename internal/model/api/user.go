package api

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserCreate struct {
	User
	Email string `json:"email" form:"email"`
	Code  string `json:"code" form:"code"`
}
