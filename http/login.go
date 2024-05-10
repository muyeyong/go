package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func sayHello(w http.ResponseWriter, r * http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprint(w, "<script>alert('xxxx')</script>")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Method)
	if r.Method == "GET" {
		timestamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write(([]byte(timestamp)))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))
		t, _ := template.ParseFiles("http/login.gtpl")
		fmt.Println(t)
		log.Println(t.Execute(w, token))
	} else {
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		fmt.Println("token", r.Form["token"])
		slice := []string{"movies", "coding", "music"}
		fmt.Println(slice)
		fmt.Println(r.Form["interest"])
		fmt.Fprint(w, r.Form["username"][0])
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