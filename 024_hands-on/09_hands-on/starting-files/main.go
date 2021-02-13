package main

import (
	"html/template"
	"log"
	"net/http"
)



var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/",home)
	http.Handle("/public/",
	http.StripPrefix("/public",
	http.FileServer(http.Dir("public"))))

	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func home(w http.ResponseWriter, r *http.Request){
	if err:=tpl.ExecuteTemplate(w,"index.gohtml",nil);err!=nil{
		log.Fatalln("Templates didn't exist.")
	}
}