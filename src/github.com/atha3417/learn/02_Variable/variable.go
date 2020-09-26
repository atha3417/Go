package main;

import ("fmt");

// Variable - Tipe Data
var namaDepan = "Atha"; // String
var namaBelakang = "Tsaqif"; // String
var namaLengkap = namaDepan + " " + namaBelakang;

func main() {
	fmt.Println("Nama depan " + namaDepan);
	fmt.Println("Perkenalkan nama saya " + namaLengkap);
	fmt.Println("Nama belakang " + namaBelakang);
}