package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Print("MASUKAN 5 ANGKA : ")
	fmt.Scan(&a, &b, &c, &d, &e)
	var hasil int
	hasil = a + b + c + d + e
	fmt.Print("HASILNYA YAITU: ", hasil)
}
