package helper

import (
	
	"fmt"
	"strings"
)

// Convert the given string to a Upper case
func ToUpperCase(str string){
	fmt.Println(strings.ToUpper(str))
}

// Convert the given string to a lower case
func ToLowerCase(str string){
	fmt.Println(strings.ToLower(str))
}

// String Search functions

// Prefix means starts with a particular string
func ToCheckHasPrefix(str string , check string){
	// HasPrefix( mainString , subString ) , checks if main string starts with the subString
	status := strings.HasPrefix(str , check)
	if status{
		fmt.Printf("%v is a prefix of %v\n",check , str)
	}else{
		fmt.Printf("%v is a not prefix of %v\n",check , str)
	}
}

// Suffix means ends with a particular string
func ToCheckHasSuffix(str string , check string){
	// HasSuffix( mainString , subString ) , checks if main string ends with the subString
	status := strings.HasSuffix(str , check)
	if status{
		fmt.Printf("%v is a suffix of %v\n",check , str)
	}else{
		fmt.Printf("%v is a not suffix of %v\n",check , str)
	}
}