package main

import "fmt"

func sumAll(number ...int) int {
	total := 0

	for _, value := range number {
		total += value
	}
	return total
}

func main (){
	total:= sumAll(10,230,24,35,1230,12,2)
	fmt.Println(total)

	slice := []int{10,23,1,23,5,123,2}  // slice parameter
	total = sumAll(slice...)
	fmt.Println(total)
}