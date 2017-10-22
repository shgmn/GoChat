package main

import (
	"log"
	"net/http"
	"sync"
	"html/template"
	"path/filepath"
)

type templateHandler struct {
	once		sync.Once
	filename	string
	// 1つのテンプレートを表す.
	templ		*template.Template
}

// HTTPリクエストを処理する.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	// ルート.
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームを開始する.
	go r.run()
	// Webサーバーを開始する.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}