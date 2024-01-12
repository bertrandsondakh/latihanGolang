package main

import "fmt"

func main()  {
	
	//var ptr *int

	//fmt.Println("what is the real value", ptr)

	mynumber := 23
	fmt.Println("the real value is", mynumber)
	var ptr = &mynumber
// *ptr artinya kayak i want to see what inside that ptr, nah how to fill inside *ptr?
// pake var ptr = &mynumber, simbol & artinya reference, jadi i want to see what inside ptr(pake simbol *), referencenya mynumber (pake simbol &)
	fmt.Println("the real value is ", &ptr)
	fmt.Println("the real value is ", *ptr)

	*ptr = *ptr+2

	fmt.Println("the real value is", mynumber)


}	