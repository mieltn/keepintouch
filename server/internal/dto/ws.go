package dto

import "github.com/gorilla/websocket"

type JoinChatReq struct {
	ChatId string `json:"chat_id"`
	UserId string `json:"user_id"`
}

type Client struct {
	ChatId string
	UserId string
	Conn *websocket.Conn
}