package dto

type Chat struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}

type ChatListReq struct {
	Limit  int
	Offset int
	Ids    []string
}

type ChatCreateReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
