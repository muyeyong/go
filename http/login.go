package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
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

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ :=	template.ParseFiles(("http/upload.gtpl"))
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handle, err := r.FormFile("uploadFile") // uploadfile
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprint(w, "%v", handle.Header)
		f, err := os.OpenFile("http/test/" + handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listenAndServe", err)
	}
}