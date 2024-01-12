package main

import (
	"bufio"
	"fmt"
	"os"
)

	func main() {
		
		welcome := "Hi, thankyou for coming, plis give us a rate for the food "
		fmt.Println(welcome)

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Name: ")
		name, _ := reader.ReadString('\n')
		fmt.Println("Hello" +name)

		fmt.Println("how was the food?")
		food, _ := reader.ReadString('\n')
		fmt.Println("Thankyou for the feedback " +name + 
		" The food is", food)

	} 	