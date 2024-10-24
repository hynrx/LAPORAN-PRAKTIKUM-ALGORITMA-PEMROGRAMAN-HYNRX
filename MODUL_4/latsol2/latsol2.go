package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukan BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukan Tinggi Badan (m): ")
	fmt.Scan(&tinggiBadan)
	beratBadan = bmi * (tinggiBadan * tinggiBadan)
	fmt.Printf("Berat badan anda: %.f", beratBadan)
}
