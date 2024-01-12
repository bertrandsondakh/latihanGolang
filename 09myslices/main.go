package main

import (
	"fmt"
	"sort"
)

func main()  {
	
	fmt.Println("welcome to slices")

	var fruitlist= []string {"apple", "tomato", "cherry"}

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "manggo", "durian")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:4])
	fmt.Println(fruitlist)

	highscores := make([]int, 4)

	highscores[0] = 200 
	highscores[1] = 400
	highscores[2] = 100
	highscores[3] = 300

	fmt.Println(highscores)

	highscores = append(highscores, 500, 600)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	var courses = []string {"hello", "hii", "how", "heyy", "what"}
	fmt.Println(courses)
	var index = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	
}