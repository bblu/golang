package main

import (
	"fmt"
	"net/http"
	"log"
)

func sayhello(res http.ResponseWriter, req *http.Request){
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form{
		fmt.Println("kv={",k,":",v,"}")
	}
	fmt.Fprintf(res, "hello world!")
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":8088",nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
