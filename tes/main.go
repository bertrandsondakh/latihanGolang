package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	
	presenttime := time.Now()
	fmt.Println(presenttime)
	fmt.Println("Current time: ", presenttime.Format("01-02-2006 15:04:05 Monday" ))

	fmt.Println("Turn your score to grade")
	fmt.Println("Please input your name: ")

	reader := bufio.NewReader(os.Stdin)

	name,_ := reader.ReadString('\n')
	fmt.Println("Hello ", name)

	
	fmt.Println("Please input your score in numbers")
	scoreInput, _ := reader.ReadString('\n')
	scoreInput = strings.TrimSpace(scoreInput)

	score, err := strconv.Atoi(scoreInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result string

	if score >= 0 && score <=50{
		result = "Youre grade is C. Dont give up " +name
	} else if score >= 51 && score <= 79{
		result = "Youre grade is B. Keep it up " +name
	} else{
		result = "Youre grade is A. Good work " +name
	}

	fmt.Println(result)
}