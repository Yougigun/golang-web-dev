package main

import (
	"html/template"
	"net/http"
	"golang-web-dev/030_sessions/02-1_session/session"
	_ "golang-web-dev/030_sessions/02-1_session/session-provider/memory"

)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
// var dbUsers = map[string]user{}      // user ID, user
// var dbSessions = map[string]string{} // session ID, user ID
var sessionManager *session.Manager
var err error

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	sessionManager,err = session.NewManager("memory","gosessionid", 3600)
	if err != nil {
		panic("Cannot New Session Manager")
	}
	go sessionManager.GC()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// process form submission
	var u user
	if req.Method == http.MethodPost {
		session := sessionManager.SessionStart(w,req)
		session.Lock()
		defer session.Unlock()
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		session.Set("username",un)
		session.Set("firstname",f)
		session.Set("lastname",l)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	exist := sessionManager.CookieExist(req)
	// get cookie
	if !exist {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}else{
		session := sessionManager.SessionStart(w,req)
		session.RLock()
		defer session.RUnlock()
		un,_ := session.Get("username")
		f,_ := session.Get("firstname")
		l,_ := session.Get("lastname")
		u := user{un.(string),f.(string),l.(string),}
		tpl.ExecuteTemplate(w, "bar.gohtml", u)
	}

}

// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV
