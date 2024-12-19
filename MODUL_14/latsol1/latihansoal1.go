package main

import "fmt"

func main() {
	var bil, j, hasil int
	fmt.Scan(&bil)
	for j = 0; j <= bil; j++ {
		if j%2 != 0 {
			hasil++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil", hasil)
}
