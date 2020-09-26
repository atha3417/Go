package main;

import ("fmt");

func main() {
	// Tipe data boolean -> valuenya hanya ada true/false

	// If Else Condition
	num := 10;
	if num > 100 {
		fmt.Println("Ya benar", num, "lebih besar dari 100");
	} else {
		fmt.Println("Salah", num, "tidak lebih besar dari 100");
	}

	color := "merah";
	if color == "hijau" {
		fmt.Println("Dor.. meletus balon hijau");
	} else {
		fmt.Println("Lanjutkan....");
	}
}