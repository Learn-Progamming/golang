package main

import "fmt"

func main(){
	var (
		ujian = 80
		absensi = 70
		lulusUjian = ujian >=  80
		lulusAbsensi = absensi >= 80
		lulus = lulusAbsensi && lulusUjian
	)	

	fmt.Println(lulus)
	fmt.Println(ujian >= 80 && absensi >= 80)
}
