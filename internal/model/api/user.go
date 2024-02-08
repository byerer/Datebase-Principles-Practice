package api

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type EmailInfo struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type RegisterInfo struct {
	User      User
	EmailInfo EmailInfo
}

type ForgetPasswordInfo struct {
	EmailInfo EmailInfo
	Password  string `json:"password"`
}
