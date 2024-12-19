package main

import "fmt"

func main() {
	var n, i int
	var prima bool
	fmt.Scan(&n)
	if n < 2 {
		fmt.Println("bukan prima")
	} else {
		prima = true
		for i = 2; i*i <= n; i++ {
			if n%i == 0 {
				prima = false
				i = n
			}
		}
		if prima {
			fmt.Println("prima")
		} else {
			fmt.Println("bukan prima")
		}
	}
}
