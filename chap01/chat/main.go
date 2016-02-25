package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`
      <html>
        <head>
	  <title>Chat チャット</title>
	</head>
        <body>
	  チャットしましょう！ / Let's chat !!
	</body>
      </html>
    `))
  })

  // Web サーバーを開始します / Start HTTP Server
  if err  := http.ListenAndServe(":8080", nil);  err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
