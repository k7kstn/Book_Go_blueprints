package main

import (
  "log"
  "net/http"
  "github.com/gorilla/websocket"
)

type room struct {
  // forward は他のクライアントに転送するためのメッセージを保持するチャネルです。
  // forward is a channel to hold message which is to be forwarded to other clients.
  forward chan []byte

  // join はチャットルームに参加しようとしているクライアントのためのチャネルです。
  // join  is a channel for clients aim to join the Chat Room.
  join chan *client

  // leave はチャットルームから退室しようとしているクライアントのためのチャネルです。
  // leave  is a channel for clients aim to leave the Chat Room.
  leave chan *client

  // clients には在室しているすべてのクライアントが保持されます。
  // clients  keep all clients in the Chat Room.
  clients map[*client]bool
}

// newRoom はすぐに利用できるチャットルームを生成して返します。
// newRoom  generates available Chat Room and returns it.
func newRoom() *room {
  return &room{
    forward: make(chan []byte),
    join: make(chan *client),
    leave: make(chan *client),
    clients: make(map[*client]bool),
  }
}

func (r *room) run() {
  for {
    select {
    case client := <-r.join:
      // 参加 Join
      r.clients[client] = true
    case client := <-r.leave:
      // 退室 Leave
      delete(r.clients, client)
      close(client.send)
    case msg := <-r.forward:
      // すべてのクライアントにメッセージを送信 Send message to All Clients.
      for client := range r.clients {
        select {
	case client.send <- msg:
	  // メッセージを送信 Message send
	default:
	  // 送信に失敗
	  delete(r.clients, client)
	  close(client.send)
	}
      }
    }
  }
}

const (
  socketBufferSize = 1024
  messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize:
    socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  socket, err := upgrader.Upgrade(w, req, nil)
  if err != nil {
    log.Fatal("ServeHTTP:", err)
    return
  }
  client := &client{
    socket: socket, 
    send: make(chan []byte, messageBufferSize),
    room: r,
  }

  r.join <- client
  defer func() { r.leave <- client }()
  go client.write()
  client.read()
}

