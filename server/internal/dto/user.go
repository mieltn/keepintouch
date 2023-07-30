package dto

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateRes struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuthRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Id           string `json:"id"`
	Username     string `json:"username"`
}

type UserRefreshReq struct {
	RefreshToken string `json:"refresh_token"`
}
