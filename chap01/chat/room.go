package main

type room struct {
  // forward は他のクライアントに転送するためのメッセージを保持するチャネルです。
  // forward is a channel to hold message which is to be forwarded to other clients.
  forward chan []byte
}

