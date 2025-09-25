package main

import "fmt"

func main() {
	var r float64
	const PI = 3.1415926535897932384626433
	fmt.Println("MENGHITUNG LUAS LINGKARAN")
	fmt.Print("Masukan jari jari: ")
	fmt.Scan(&r)
	var hasil float64 = r * r * PI
	fmt.Println("Hasilnya adalah:", hasil)

}
