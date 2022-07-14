package main

import "fmt"

func main (){
	name:= "Stevanus"
	
	if name == "Cristian" {
		fmt.Println("Hallo Cristian")
	} else if name == "Stevanus" {
		fmt.Println("Hallo Stevanus")
	} else {
		fmt.Println("Hi",name)
	}
	
	// length:= len(name)
	// if length > 5 {
	// 	fmt.Println("nama terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	if length := len(name) ;length > 5 { // with short statement
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}


}