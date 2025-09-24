package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Masukkan input: ")
	fmt.Scanln(&input)
	
	if input == "66" || input == "97" || input == "103" || input == "117" || input == "115" {
		fmt.Println("Bagus")
	} else if strings.ToUpper(input) == "SNO" {
		fmt.Println("TOP")
	} else {
		fmt.Println("Input tidak dikenal")
	}
}
