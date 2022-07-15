package main

import "fmt"

func getFullName2()(firstName, middleName, lastName string){
	firstName = "Cristian"
	middleName = "Stevanus"
	lastName = "Oloan"
	return
}
func main (){
	firstName, middleName, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}