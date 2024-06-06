package main

import (
	"fmt"
	"net/http"

	"github.com/kokim2020/web_hello/number_types"
)

/*
All Go programs start running from a function called `main` in a package called `main`
*/
func main() {
	number_types.PrintHello()
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}
