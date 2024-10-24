package main

import "fmt"

func main() {
	var totalBelanja, diskon int
	fmt.Print("Masukkan total harga belanja: ")
	fmt.Scanln(&totalBelanja)
	fmt.Print("Masukkan diskon: ")
	fmt.Scanln(&diskon)

	totalDiskon := totalBelanja - (totalBelanja * diskon / 100)
	fmt.Print("Total harga belanja setelah diskon: ", totalDiskon)
}
