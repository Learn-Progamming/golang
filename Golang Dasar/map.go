package main

import "fmt"

func main(){
	person:= map[string]string{
		"name": "cris",
		"address": "cimahi",
	}

	person["title"] = "programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])


	book:= make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}