package main

import "fmt"

func main(){
	var name1 = "Cristian"
	var name2 = "Cristian"

	var result bool = name1 == name2
	fmt.Println(result)

	var name3 = "Cristian"
	var name4 = "Stevanus"

	var result2 bool = name3 == name4
	fmt.Println(result2)

	var result3 bool = name3 > name4
	fmt.Println(result3)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
