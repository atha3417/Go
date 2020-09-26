package main;

import ("fmt");

// Menulis 2 variable dalam satu baris
var namaDepan, namaBelakang = "Atha", "Tsaqif";
var namaLengkap = namaDepan + " " + namaBelakang;

// Menulis variable dengan nilai kosong
// var sekolah = ""; // cara pertama
var sekolah string; // cara kedua

func main() {
	fmt.Println("Perkenalkan nama saya " + namaLengkap);
	fmt.Println("Bersekolah di " + sekolah);
}