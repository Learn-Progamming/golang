package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You're Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
// func blacklist1(name string) bool{
// 	return name == "admin"
// }
// func blacklist1(name string) bool{
// 	return name == "admin"
// }
func main (){
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin",blacklist)
	registerUser("root",blacklist)

	registerUser("root", func(name string) bool{
		return name == "root"
	})
	registerUser("ega", func(name string) bool{
		return name == "root"
	})
}