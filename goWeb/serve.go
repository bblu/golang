package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
)

func sayHello(res http.ResponseWriter, req *http.Request){
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
func muxHello(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "hello world!")
}
func muxEcho(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path + " from serveMux")
}
//if you visit / you will get '404 page not found' error
func main() {
	http.HandleFunc("/", sayHello)
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", muxEcho)
	mux.HandleFunc("/hello", muxHello)

	err := http.ListenAndServe(":8088",mux)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
	log.Print("serve listen :8088")

}
