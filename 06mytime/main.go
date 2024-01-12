package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Golang time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	
	createdDate := time.Date(2023, time.December, 18, 15, 15, 15, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))


}