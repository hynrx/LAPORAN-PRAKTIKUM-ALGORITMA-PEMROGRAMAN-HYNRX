package main

import "fmt"

func main() {
	var (
		nama, nim, kelas string
	)
	fmt.Print("Masukan Nama, NIM, dan Kelas : ")
	fmt.Scan(&nama, &nim, &kelas)
	fmt.Print("Perkenalkan saya adalah ", nama, ", salah satu mahasiswa Prodi S1-IF dari kelas ", kelas, " dengan nim ", nim, ".")
}
