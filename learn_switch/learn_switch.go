package learn_switch

import (
	"fmt"
	"os"
)

func LearnSwitch() {
	word := os.Args[0]
	switch word {
	case "hello":
		fmt.Println("hi your self")
	case "good bye":
		fmt.Println("so long")
	case "greetings":
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said")
	}
}
