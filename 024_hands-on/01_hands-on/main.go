package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogImg)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<h1>This is from dog</h1>
						<img src="/dog.jpg"/>`)
}

func dogImg(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./dog.jpg")
}