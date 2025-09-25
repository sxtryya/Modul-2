package main

import "fmt"

func main() {
	var nama string
	var nim int64
	var kelas string
	fmt.Print("Masukan Nama, NIM, dan Kelasmu: ")
	fmt.Scan(&nama, &nim, &kelas)
	fmt.Print("Halo Semua, Perkenalkan nama saya ", nama)
	fmt.Print(" saya merupakan mahasiswa Prodi S1-IF dari kelas ", kelas, " dengan NIM: ", nim)
}
