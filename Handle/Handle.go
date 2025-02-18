package main

import "net/http"

type MyHandler struct {
}

func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my first handler"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testHandler"))
}
func main() {

	mh := MyHandler{}
	http.Handle("/myhandler", mh)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is my handlefunc"))
	})
	http.HandleFunc("/welcome", welcome)
	http.Handle("/testHandler", http.HandlerFunc(testHandler))
	http.ListenAndServe("localhost:4070", nil)

}
