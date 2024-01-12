package main

import "fmt"

func main()  {
	
	var fruitlist [4]string

	fruitlist [0] = "apple"
	fruitlist [1] = "orange"
	fruitlist [2] = "durian"
	fruitlist [3] = "watermelon"

	fmt.Println("fruitlist: ", fruitlist)
	fmt.Println("fruitlist: ", len(fruitlist))

	var vegilist = [3]string{"kangkung", "mushroom", "carrot"}
	fmt.Println("vegilist: ", vegilist)	
	fmt.Println("vegilist: ", len(vegilist))	
}