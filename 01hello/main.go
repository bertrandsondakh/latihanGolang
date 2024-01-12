package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	
	fmt.Println("Welcome to Pizza Restaurant customer service")
	fmt.Println("Please input your firstname: ")

	reader := bufio.NewReader(os.Stdin)

	firstname,_ := reader.ReadString('\n')
	fmt.Println("Hello", firstname)

	fmt.Println("Please input your lastname: ")
	lastname,_ := reader.ReadString('\n')
	fmt.Println("Hello Mr.", lastname)

	fmt.Println("From scale 1 to 10, rate our customer services: ")
	rate,_ := reader.ReadString('\n')
	fmt.Println("Thank you for the rating", rate,"Mr.", lastname)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rate), 64)
	if err != nil {
		fmt.Println("erorr")
	}	else{
		fmt.Println("Added one to your rating", numRating+1) 
		}


}