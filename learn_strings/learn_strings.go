package learn_strings

import "fmt"

func LearnStrings() {
	// assign string in the variable
	s := `Hello,"DEMO" 
						how are u?`
	s2 := s[2:4]
	// assign array into the numbers variable
	numbers := [...]int{1, 2, 3, 4, 5}
	// You can use runes to work with Unicode characters.
	// For example, you can iterate over the characters in a string,
	// which might contain multi-byte UTF-8 encoded characters,
	// and work with each character as a rune
	r3 := rune(0x1F600)
	fmt.Println(numbers)
	fmt.Println(s2)
	// print unicode rune character
	fmt.Println(string(r3))
	// Rune Literal
	// A rune literal is just a character enclosed in single quotes.
	// You can also use Unicode escape sequences to represent runes:
	r4 := '\u4e16'
	fmt.Println(string(r4))
}
