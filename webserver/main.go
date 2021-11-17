package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1> Hello, I am Pizza!<h1>")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}

//手动完成的
//1.git pull
//2.git push -> git pull
//3.deploy