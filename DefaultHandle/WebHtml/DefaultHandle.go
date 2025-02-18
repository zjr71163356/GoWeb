package main

import "net/http"

type MyStruct struct{}

func (ms *MyStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("我的struct"))
}
func main() {
	fs := http.FileServer(http.Dir("WebHtml"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.HandleFunc("/web/index", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/web/index.html", http.StatusFound)
	})
	http.ListenAndServe("localhost:4070", nil)
}
