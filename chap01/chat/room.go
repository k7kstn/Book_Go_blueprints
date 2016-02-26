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

