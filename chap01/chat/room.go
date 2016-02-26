package main

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

