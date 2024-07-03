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

// Let's take a specific iteration of the loop as an example:

// If k is 0 and v is 72, this means the character at index 0 of the string is H.
// k: The index of the character.

// For example, k = 0 for the first character 'H', k = 1 for the second character 'e', and so on.
// v: The Unicode code point of the character.

// For 'H', the Unicode code point is 72.
// string(v): Converts the Unicode code point back to the character.

// string(72) returns 'H'.
func LearnForRangeLoop() {
	s := "Hello world"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
