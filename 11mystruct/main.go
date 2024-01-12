package main

import "fmt"

func main()  {
	bertrand := User{"bertrand", "bertrandsondakh@gmail.com", true, 22}
	fmt.Println(bertrand)
	fmt.Printf("bertrand details are %+v \n", bertrand)
	fmt.Printf("bertrand email %v and age is %v", bertrand.Email, bertrand.Age)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int

}