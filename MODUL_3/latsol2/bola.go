package main

import "fmt"

func main() {
	var j, v, l float64
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&j)
	v = (4.0 / 3.0) * 3.1415926535 * (j * j * j)
	l = 4 * 3.1415926535 * (j * j)
	fmt.Print("Bola dengan jari-jari ", j, "memiliki volume ", v, " dan luas kulit ", l)
}
