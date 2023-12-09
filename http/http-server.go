package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("/hello receive conn ", req.Host)
	fmt.Fprintf(w, "world!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("/headers receive conn ", req.Host)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":9999", nil)
}
