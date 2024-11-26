package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan : ")
	fmt.Scanln(&b)
	fmt.Print("Faktor : ")
	for j := 1; j <= b; j++ {
		if b%j == 0 {
			fmt.Print(j, " ")
		}
	}
	fmt.Print("\n")
	if b%2 == 0 || b%3 == 0 || b%5 == 0 || b%7 == 0 && b != 1 && b != 2 && b != 3 && b != 5 && b != 7 {
		fmt.Println("Prima : false")
	} else {
		fmt.Println("Prima : true")
	}
}
