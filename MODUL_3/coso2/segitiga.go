package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Print("Masukan alas: ")
	fmt.Scan(&alas)
	fmt.Print("Masukan tinggi: ")
	fmt.Scan(&tinggi)
	luas = (alas * tinggi / 2)
	fmt.Println("Hasilnya adalah:", luas)
}
