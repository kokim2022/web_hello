package learn_if_for

import "fmt"

func LearnIf() {
	a := 10
	// initialization within the if statement
	if b := a / 2; b > 5 {
		fmt.Println("a is greater than ")
	} else {
		fmt.Println("a is less than or equal to 5:", a, b)
	}
}

func LearnFor() {
	for i := 1; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}
}
