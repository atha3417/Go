package main;

import ("fmt");

var nama = "Atha";

func printNumber() /* tipe data yang ingin di kembalikan */ int {
	fmt.Println("Kamu memanggil fungsi printNumber");
	// return
	return 10;
}

func multiply(/* parameter */ angka1 int, angka2 int) int {
	return angka1 * angka2;
}

func main() {
	fmt.Println("Fungsi printNumber menghasilkan...", printNumber());
	fmt.Println("Fungsi multiply menghasilkan...", multiply(10, 5));
}