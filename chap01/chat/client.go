package main

import (
  "github.com/gorilla/websocket"
)

// client  represents 1 user who is chatting.
// client  はチャットを行っている１人のユーザーを表します。
type client struct {
  // socket はこのクライアントのための WebSocket です。
  // socket  is a WebSocket  for this client.
  socket  *websocket.Conn

  // send はメッセージが送られるチャネルです / is a channel to send messages.
  send chan []byte

  // room はこのクライアントが参加しているチャットルームです。
  // room is a Chat Room which this client is participating.
  room  *room
}

