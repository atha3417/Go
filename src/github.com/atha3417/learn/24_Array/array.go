package main;

import ("fmt");

func main() {
	// Array //! Harus ditentukan jumlah item
	var members [3] string;
	//! Index array dimulai dari 0
	// Mengisi array

	// Cara 1
	members[0] = "Atha";
	members[1] = "Nafisa";
	members[2] = "Fathan";

	// Cara 2
	numbers := [4] int {11232, 208, 32334, 423};

	// Mengubah isi array
	members[0] = "Aiman";
	numbers[0] = 323232;

	// Jika kita tidak tahu berapa jumlah item di array
	// kita bisa gunakan
	numbers2 := [...] int {12121, 23244, 3434, 23232};

	fmt.Println(members);
	fmt.Println(numbers);
	fmt.Println(numbers2);
}