package main

import "fmt"

func main(){
	var names[3]string
	names[0] = "Cristian"
	names[1] = "Stevanus"
	names[2] = "Oloan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		95,
		80,
	}
	fmt.Println(values)
	fmt.Println(len(names))
	fmt.Println(len(values))
}