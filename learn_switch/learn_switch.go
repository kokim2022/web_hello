package learn_switch

import (
	"fmt"
	"os"
)

// os.Args[0] used to get input from the cli
func LearnSwitch() {
	// In Go, os.Args is a slice of strings that contains the command-line arguments passed to the program.
	//  The first element, os.Args[0], is the name of the program itself. Subsequent elements,
	//   os.Args[1], os.Args[2], etc., are the actual arguments passed to the program.
	word := os.Args[1]
	switch word {
	case "hi":
		fmt.Println("it is very informal")
		// 		Key Points
		// 1. Default Behavior: In Go, the switch statement breaks after each case by default.
		// 2. Explicit Fallthrough: To allow execution to proceed to the next case, you must explicitly use the fallthrough keyword.
		// 3. Conditions Ignored: When fallthrough is used, the next case's condition is not checked.
		fallthrough
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
