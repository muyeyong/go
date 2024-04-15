package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r * http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprint(w, "hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("http/login.gtpl")
		fmt.Println(t)
		log.Println(t.Execute(w, nil))
	} else {

	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listenAndServe", err)
	}
}