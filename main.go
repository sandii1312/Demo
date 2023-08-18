package main

import (
	"bufio"
	"fmt"
	"go-fiber/helper"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var sub string
	fmt.Print("Enter the main String : ")
	main, _ := reader.ReadString('\n')
	fmt.Print("Enter the sub String : ")
	fmt.Scan(&sub)
	fmt.Printf("The index of substring %v is %v in main string %v\n", sub, helper.ToCheckIndexOfElement(main, sub), main)
}
