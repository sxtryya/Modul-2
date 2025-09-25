package main

import "fmt"

func main() {
	var suhu int64
	fmt.Println("KONVERSI SUHU DARI FAHREINHEIT (F) KE CELCIUS (C)")
	fmt.Print("Masukan suhu dalam bentuk fahreinheit: ")
	fmt.Scan(&suhu)
	var C = (suhu - 32) * 5 / 9
	fmt.Println("Hasilnya adalah:", C, "°C")

}
