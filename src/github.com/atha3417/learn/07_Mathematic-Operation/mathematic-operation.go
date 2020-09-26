package main;

import ("fmt");

func main() {
	// menulis variable kosong
	var hasil int;

	// Operasi matematika (+, -, /, *, %)
	a := 10;
	b := 20;
	c := 5;
	hasil = b * a / c;
	
	fmt.Println("Hasilnya...", hasil);

	// Increment & Decrement (++, --)
	a++;
	b--;

	fmt.Println("Hasilnya...", a, "dan b adalah", b);
}