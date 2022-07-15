package main

import "fmt"

func getFullName() (string,string){
	return "Cristian","Stevanus"
}
func main(){
	// firstName, lastName := getFullName()
	firstName, _ := getFullName() // if u want ignore one of all return value, u can replace variable with underscore _
	fmt.Println(firstName)

}