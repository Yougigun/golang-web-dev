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
	http.Handle("/resources/",
	http.StripPrefix("/resources",
	http.FileServer(http.Dir("public"))))

	http.HandleFunc("/",home)

	log.Fatalln(http.ListenAndServe(":8080",nil)) 
}

func home(w http.ResponseWriter, r *http.Request){
	err := tpl.ExecuteTemplate(w,"index.gohtml",nil)
	if err !=nil{
		log.Fatalln("Template didn't find.")
	}
}