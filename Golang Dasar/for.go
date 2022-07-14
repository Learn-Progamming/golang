package main

import "fmt"

func main(){
	// counter:= 1

	// for counter <= 10{					// looping without statement (simple)
	// 	fmt.Println("perulangan ke",counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++{  // looping with init statement and post statement
	// 	fmt.Println("perluangan ke", counter)
	// }

	slice := []string{"Cristian","Stevanus","Oloan","Fajar","Maulana"}  // looping for with slicing array
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	for i, value := range slice{ // looping for range with index
		fmt.Println("index",i,"=",value)
	}

	for _, value := range slice{ // looping for range without index, show value only
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Cristian"
	person["title"] = "Programmer"

	for _, value := range person{ // looping for range to show value and key on map
		fmt.Println(value)
	}
}