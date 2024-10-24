package main

import "fmt"

func main() {
	var x int
	var phi, y float64
	fmt.Scan(&x)
	phi = 3.14
	y = phi * float64(x*x)
	fmt.Print(y)
}
