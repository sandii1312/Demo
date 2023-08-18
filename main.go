package main

import (
	"bufio"
	"os"
	"go-fiber/helper"
	"fmt"
	"log"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	var sub string
	fmt.Print("Enter the main String : ")
	main , err := reader.ReadString('\n') 	
	fmt.Print("Enter the sub String : ")
	fmt.Scan(&sub)
	if err != nil{
		log.Fatal(err)
	}else{
		helper.ToCheckHasSuffix(strings.TrimSpace(main) , sub)
	}
	
}