// package main

// import (
// 	"fmt"
// 	"io"
// 	// "net/http"
// 	"runtime"

// 	"golang.org/x/net/websocket"
// )

// type Server struct {
// 	conns map[*websocket.Conn]bool
// }

// func NewServer() *Server {
// 	return &Server{
// 		conns: make(map[*websocket.Conn]bool),
// 	}
// }

// func (s *Server) handleWS(ws *websocket.Conn) {
// 	fmt.Println(
// 		"new incoming connection for client: ",
// 		ws.RemoteAddr(),
// 	)

// 	s.conns[ws] = true

// 	s.readLoop(ws)
// }

// func (s *Server) readLoop(ws *websocket.Conn) {
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := ws.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			fmt.Println("reading error: ", err)
// 			continue
// 		}
// 		msg := buf[:n]
// 		fmt.Println(string(msg))
// 		s.broadcast(msg)
// 	}
// }

// func (s *Server) broadcast(b []byte) {
// 	for ws := range s.conns {
// 		go func(ws *websocket.Conn) {
// 			if _, err := ws.Write(b); err != nil {
// 				fmt.Println("write error: ", err)
// 			}
// 		}(ws)
// 	}
// }

// func main() {
// 	fmt.Println(runtime.NumCPU())
// 	// server := NewServer()
// 	// http.Handle("/ws", websocket.Handler(server.handleWS))
// 	// http.ListenAndServe(":3000", nil)
// }