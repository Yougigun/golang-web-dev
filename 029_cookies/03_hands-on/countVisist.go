package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", count)
	// http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func count(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("count")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "count",
			Value: "1",
			Path:  "/",
		})
		fmt.Fprintln(w, "You have visited. 1 times.")
		return
	}
	cInt, err := strconv.ParseInt(c.Value, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
	cInt++
	http.SetCookie(w, &http.Cookie{
		Name:  "count",
		Value: fmt.Sprintf("%v", cInt),
		Path:  "/",
	})
	fmt.Fprintf(w, "You have visited. %v times.", cInt)
}

// Using cookies, track how many times a user has been to your website domain.
