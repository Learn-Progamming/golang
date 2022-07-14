package main

import "fmt"

func main (){
	name:= "Fajar"

	switch name {
	case "Cristian":
		fmt.Println("Hallo Cristian")
		fmt.Println("Hallo Cristian")
	case "Stevanus":
		fmt.Println("Hallo Stevanus")
		fmt.Println("Hallo Stevanus")
	default:
		fmt.Println("Hi",name)
		fmt.Println("Hi",name) 
	}

	switch length := len(name) ; length > 5{ // with short statement
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("name benar")
	}


length:= len(name)
	switch { //switch case simple
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama cukup panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}