package main

import (
  "log"
  "net/http"
  "text/template"
  "path/filepath"
  "sync"
)

// templ は１つのテンプレートを表します templ indicates 1 template
type templateHandler struct {
  once     sync.Once
  filename string
  templ    *template.Template
}

// ServeHTTP は HTTPリクエストを処理します handles HTTP requests
func  (t *templateHandler)  ServeHTTP(w http.ResponseWriter, r *http.Request) { 
  t.once.Do(func()  {
    t.templ = 
        template.Must(template.ParseFiles(filepath.Join("templates",
	    t.filename)))
  })
  t.templ.Execute(w, nil)  // return value should be checked.
}

func main() {
  // Root
  http.Handle("/", &templateHandler{filename: "chat.html"})

  // Web サーバーを開始します / Start HTTP Server
  if err  := http.ListenAndServe(":8080", nil);  err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
