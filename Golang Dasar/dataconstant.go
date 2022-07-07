package main

import "fmt"

func main(){
	const firstName string = "Cristian"
	const lastName = "Stevanus"

	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		satuName string = "Cristian"
		duaName 				= "Stevanus"
		value_baru			= 1000	 
	)
	
	fmt.Println(satuName)
	fmt.Println(duaName)
	fmt.Println(value_baru)
}