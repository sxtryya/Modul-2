package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Print("Masukan angka angka penjumlahan: ")
	fmt.Scan(&a, &b, &c, &d, &e)
	var hasil = a + b + c + d + e
	fmt.Println("Jadi hasil penjumlahan", a, "+", b, "+", c, "+", d, "+", e, "adalah: ", hasil)

}
