package main

import "fmt"

/*
All Go programs start running from a function called `main` in a package called `main`
*/
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

// pass function as a variable
//
//	&
//	 its return type is also a function
func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func main() {
	// number_types.PrintInteger()
	// learn_strings.LearnStrings()
	// learn_switch.LearnSwitchWithBoolean()
	// a := learn_funcs.AddNumbers(1, 2)
	// fmt.Println(a)
	// a = learn_funcs.AddNumbers(3, 4)
	// fmt.Println(a)
	// a = learn_funcs.AddNumbers(3, 5)
	// fmt.Println(a)
	// http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
	// 	name := req.URL.Query().Get("name")
	// 	rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	// })
	// http.ListenAndServe(":8080", nil)

	// When a function returns multiple values,
	//  but you are interested in only some of them,
	// you can use _ to ignore the values you don't need.
	// div, remainder := learn_funcs.DivAndRemainder(2, 3)
	// fmt.Println(div, remainder)

	// div, _ = learn_funcs.DivAndRemainder(10, 2)
	// fmt.Println(div)

	// _, remainder = learn_funcs.DivAndRemainder(100, -100)
	// fmt.Println(remainder)

	// myAddOne := learn_more_funcs.AddOne
	// fmt.Println(myAddOne(3))
	// fmt.Println(learn_more_funcs.AddOne(3))

	b := 2
	myAddOne := func(a int) {
		b = a + b
	}
	myAddOne(6)
	fmt.Println(b)

	// call first function
	addOne := makeAdder(1)

	// pass function as parameter
	doubleAddone := makeDoubler(addOne)
	fmt.Println(doubleAddone(1))
	// addTwo := makeAdder(2)
	// call second function
	// fmt.Println(addOne(1))
	// fmt.Println(addTwo(1))

	// learn_more_funcs.LearnMoreFunctions()
}
