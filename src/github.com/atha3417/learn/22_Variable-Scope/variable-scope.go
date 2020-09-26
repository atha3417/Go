// Variable Scope
package main;

import ("fmt");

var bonus = 50; // variable global, karena ada //! di luar fungsi

func hitungTotal(gaji int) int {
	bonus = 100; // bonus akan //! berubah
	return gaji + bonus;
}

func main() {
	// var bonus = 50; // local variable, karena ada di //! dalam fungsi
	fmt.Println(hitungTotal(10));
	fmt.Println("Bonus =", bonus);
}