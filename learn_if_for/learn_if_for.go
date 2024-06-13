package learn_if_for

import "fmt"

func LearnIfFor() {
	a := 10
	if b := a / 2; b > 5 {
		fmt.Println("a is greater than ")
	} else {
		fmt.Println("a is less than or equal to 5:", a, b)
	}
}
