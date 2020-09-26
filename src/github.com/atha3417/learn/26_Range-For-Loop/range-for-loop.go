package main;

import ("fmt");

func main() {
	days := [7] string {"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"};

	// Range
	for _, day := range days {
		// fmt.Printf("Hari %d = %s\n", i, day); //! Gunakan Printf untuk menggunakan (%)
		fmt.Printf("Hari %s\n", day);
	}
}