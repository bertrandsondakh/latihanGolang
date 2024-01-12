package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("welcome to switch case golang")

	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6)+1
	fmt.Println("your value is ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("dice value is 1, you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot, and dice again")
	default: 
		fmt.Println("what was that!")
	}

}  
