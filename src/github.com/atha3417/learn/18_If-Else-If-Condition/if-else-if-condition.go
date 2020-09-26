package main;

import ("fmt");

func main() {
	utang := 5000;
	uang := 1000;

	if uang >= utang {
		fmt.Println("Utang lunas... tunggu kembaliannya");
	} else if uang == utang {
		fmt.Println("Uangnya pass... lunas");
	} else {
		fmt.Println("Uang kamu tidak cukup");
	}
}