package dto

import "time"

type Message struct {
	Id        string    `json:"id"`
	ChatId    string    `json:"chat_id"`
	UserId    string    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageByChatIdReq struct {
	Limit  int
	Offset int
	Id     string
}

type MessageCreateReq struct {
	ChatId string `json:"chat_id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}
