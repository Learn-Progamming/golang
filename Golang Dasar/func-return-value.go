package main

import "fmt"

func getHello(name string)string{
	// result := "hello"
	// return result

	if name == ""{
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}
func main (){
	result := getHello("Cristian")
	fmt.Println(result)

	fmt.Println(getHello(""))
}