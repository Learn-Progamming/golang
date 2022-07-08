package main

import "fmt"

func main(){
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
	  "November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // Panjang Data Slicer
	fmt.Println(cap(slice1)) // Panjang Kapasitas
	months[5] = "Diubah" // Mengubah Data Array
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Cris")
	fmt.Println(slice3)

	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)


	// New Slice with Array
	newSlice:= make ([]string, 2, 5)
	newSlice[0] = "Cristian"
	newSlice[1] = "Stevanus"
	fmt.Println(newSlice)

	// Copy Slice
	copySclie:= make([]string, len(newSlice), cap(newSlice))
	copy(copySclie, newSlice)
	fmt.Println(copySclie)

	// Perbedaan Array vs Slice
	iniArray:= [...]int{1,2,3,4,5,6}
	iniSlice:= []int{1,2,3,4,5,6}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}