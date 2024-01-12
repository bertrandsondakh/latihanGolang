package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	
	fmt.Println("Welcome")
	fmt.Println("Please rate our pizza from 1 to 10")

	reader := bufio.NewReader(os.Stdin)

	rate,_ := reader.ReadString('\n')
	fmt.Println("Thankyou for rate: ", rate)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rate), 64)

	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("Add one to your rating: ", numRating+1)
	}
}