package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// req.Form includes the url query parameters and  form data(in body)
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	fmt.Fprintf(w,"req.Form:%v<br/>",req.Form)
	fmt.Fprintf(w,"req.Form:%v<br/>",req.PostForm)
	fmt.Fprintf(w,"req.URL:%v<br/>",req.URL)
	fmt.Fprintf(w,"req.Host:%v<br/>",req.Host)
	fmt.Fprintf(w,"req.Method:%v<br/>",req.Method)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
