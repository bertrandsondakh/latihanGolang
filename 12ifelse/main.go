package main

import "fmt"

func main()  {
	
	fmt.Println("welcome to if else golang")

	score := 85
	var result string

	if score <= 50 {
		result = "C"
	} else if score >= 51 && score <=79 {
		result = "B"
	} else {
		result = "A"
	}
	fmt.Println("your current grade is", result)
	

	// odd or even number

	if 9%2 == 0{
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
}