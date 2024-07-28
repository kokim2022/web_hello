package learn_more_funcs

import "fmt"

func AddOne(a int) int {
	return a + 1
}

func AddTwo(a int) int {
	return a + 2
}

func PrintOperation(a int, f func(a int) int) {
	fmt.Println(f(a))
}

func LearnMoreFunctions() {
	PrintOperation(1, AddOne)
	PrintOperation(1, AddTwo)
}
