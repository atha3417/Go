package main;

import ("fmt");

func main() {
	// Menghitung angka deret dengan for loop
	total := 0;
	for i := 0; i <= 10; i++ {
		total = total + i;
	}
	fmt.Println("Hasilnya adalah...", total);
}