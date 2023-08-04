package broadcaster

import (
	"context"

	"github.com/mieltn/keepintouch/internal/dto"
)

type Service struct {
	activeConnections map[string][]*dto.Client
	broadcast chan *dto.Message
	register chan *dto.Client
	unregister chan *dto.Client
}

func NewService() *Service {
	return &Service{
		activeConnections: make(map[string][]*dto.Client),
		broadcast: make(chan *dto.Message),
		register: make(chan *dto.Client),
		unregister: make(chan *dto.Client),
	}
}

func (s *Service) Run() {
	for {
		select {
		case cl := <-s.register:
			if _, ok := s.activeConnections[cl.ChatId]; ok {
				if !s.hasConnection(cl) {
					s.activeConnections[cl.ChatId] = append(s.activeConnections[cl.ChatId], cl)
				}
			}

		case cl := <-s.unregister:
			if _, ok := s.activeConnections[cl.ChatId]; ok {
				s.activeConnections[cl.ChatId] = s.updateConnections(cl)
			}

		case msg := <-s.broadcast:
			if _, ok := s.activeConnections[msg.ChatId]; ok {
				for _, client := range s.activeConnections[msg.ChatId] {
					if err := client.Conn.WriteJSON(msg); err != nil {
						client.Conn.Close()
						s.unregister <- client
					}
				}
			}
		}
	}
}

func (s *Service) GetBroadcast() chan *dto.Message {
	return s.broadcast
}

func (s *Service) GetRegister() chan *dto.Client {
	return s.register
}

func (s *Service) GetUnregister() chan *dto.Client {
	return s.unregister
}

func (s *Service) hasConnection(client *dto.Client) bool {
	for _, activeConn := range s.activeConnections[client.ChatId] {
		if client.UserId == activeConn.UserId {
			return true
		}
	}
	return false
}

func (s *Service) updateConnections(client *dto.Client) []*dto.Client {
	var connsUpdated []*dto.Client
	for _, activeConn := range s.activeConnections[client.ChatId] {
		if client.UserId != activeConn.UserId {
			connsUpdated = append(connsUpdated, client)
		}
	}
	return connsUpdated
}

func (s *Service) Close(ctx context.Context) error {
	close(s.broadcast)
	close(s.register)
	close(s.unregister)
	return nil
}