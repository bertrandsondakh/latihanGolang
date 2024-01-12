package main

import "fmt"

func main()  {
	
	languages := make(map[string]string)

	languages["ab"] = "abc"
	languages["de"] = "def"
	languages["gh"] = "gh"

	fmt.Println("list of alphabets", languages)
	fmt.Println("ab means", languages["ab"])

	delete(languages, "gh")
	fmt.Println(languages)

	//loops

	for key, value := range languages{
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}