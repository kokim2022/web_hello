package main

import (
	"github.com/kokim2020/web_hello/learn_switch"
)

/*
All Go programs start running from a function called `main` in a package called `main`
*/
func main() {
	// number_types.PrintInteger()
	// learn_strings.LearnStrings()
	learn_switch.LearnSwitchWithBoolean()
	// http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
	// 	name := req.URL.Query().Get("name")
	// 	rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	// })
	// http.ListenAndServe(":8080", nil)
}
