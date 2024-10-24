package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64

	fmt.Print("Masukkan titik A: ")
	fmt.Scan(&ax, &ay)

	fmt.Print("Masukkan titik B: ")
	fmt.Scan(&bx, &by)

	fmt.Print("Masukkan titik C: ")
	fmt.Scan(&cx, &cy)

	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(cx-ax, 2) + math.Pow(cy-ay, 2))

	hasil := math.Max(ab, math.Max(bc, ca))

	fmt.Print("Panjang sisi terpanjang: ", hasil)
}
