package main

import "fmt"

func main() {
	var t int
	var k bool
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&t)
	k = (t%400 == 0) || (t%4 == 0 && t%100 != 0)
	fmt.Println(k)
}
