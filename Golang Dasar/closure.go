package main 

import "fmt"


func main (){
	name := "Cristian"
	counter := 0

	increment := func(){
		name := "Budi"
		fmt.Println("Increment")
		fmt.Println(name)
		counter ++
	}

	increment()
	increment()
	
	fmt.Println(counter)
	fmt.Println(name)
}