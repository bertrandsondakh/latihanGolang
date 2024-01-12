package main

import "fmt"

func main()  {
	
	fmt.Println("welcome to loops in golang")

	days := []string{"monday", "thuesday", "wedenesday", "thursday", "friday"}
	fmt.Println(days)

	//for i:=0; i <len(days); i++{
	//	fmt.Println(days[i])
	//} 

	//for i:= range days{
	//	fmt.Println(days[i])
	//}

	rouguevalue := 1

	for rouguevalue < 10{

		if rouguevalue ==2{
			goto lco
		}

		if rouguevalue ==5{
			rouguevalue++
			continue
		}

		fmt.Println("the value is", rouguevalue)
		rouguevalue++
	}

	lco:
	fmt.Println("Jump at welcome to macos")
}
