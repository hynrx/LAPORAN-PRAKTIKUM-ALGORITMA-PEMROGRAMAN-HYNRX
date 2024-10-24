package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scan(&celsius)
	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273.15
	fmt.Println("Suhu dalam derajat Reamur:", reamur)
	fmt.Println("Suhu dalam derajat Fahrenheit:", fahrenheit)
	fmt.Println("Suhu dalam derajat Kelvin:", kelvin)
}
