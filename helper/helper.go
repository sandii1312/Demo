package helper

import (
	"fmt"
	// redefined the imported string package name as s
	s "strings"
)

// Convert the given string to a Upper case
func ToUpperCase(str string) {
	fmt.Println(s.ToUpper(str))
}

// Convert the given string to a lower case
func ToLowerCase(str string) {
	fmt.Println(s.ToLower(str))
}

// String Search functions

// Prefix means starts with a particular string
func ToCheckHasPrefix(str string, check string) {
	// HasPrefix( mainString , subString ) , checks if main string starts with the subString
	status := s.HasPrefix(str, check)
	if status {
		fmt.Printf("%v is a prefix of %v\n", check, str)
	} else {
		fmt.Printf("%v is a not prefix of %v\n", check, str)
	}
}

// Suffix means ends with a particular string
func ToCheckHasSuffix(str string, check string) {
	// HasSuffix( mainString , subString ) , checks if main string ends with the subString
	status := s.HasSuffix(str, check)
	if status {
		fmt.Printf("%v is a suffix of %v\n", check, str)
	} else {
		fmt.Printf("%v is a not suffix of %v\n", check, str)
	}
}

// Index(str , subStr) returns the index of the particular substring
func ToCheckIndexOfElement(str string, subStr string) int {
	return s.Index(str, subStr)
}
