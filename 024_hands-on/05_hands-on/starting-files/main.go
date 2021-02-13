package main

import (
	"html/template"
	"log"
	"net/http"
)


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs:=http.FileServer(http.Dir("public"))
	http.Handle("/pics",fs)
	http.HandleFunc("/",home)
	http.ListenAndServe(":8080",nil)
}

func home(w http.ResponseWriter, r *http.Request){
	log.Println(r.RemoteAddr,r.URL.Path)
	if err:=tpl.Execute(w,nil);err!=nil{
		log.Fatalln("template didn't execute",err)
	}
}

