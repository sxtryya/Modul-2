package main

import "fmt"

func main() {
	var x float64
	fmt.Println("persamaan f(x)= 2/(x+5) + 5")
	fmt.Print("Masukan angka pada variabel x: ")
	fmt.Scan(&x)
	var hasil = 2/(x+5) + 5
	fmt.Println("hasilnya adalah", hasil)
}
