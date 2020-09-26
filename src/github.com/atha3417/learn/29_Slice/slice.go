package main;

import ("fmt");

func main() {
	member := [5] string {"Atha", "Nafisa", "Fathan", "Aiman", "Fayza"};

	// Membuat slice dari array yang sudah ada (semua data)
	memberSlice := member[:];

	// Membuat slice dari array yang sudah ada (dari pertama ke terakhir)
	memberSlice2 := member[1:4];

	// Membuat slice dari array yang sudah ada (dari 1 ke terakhir)
	memberSlice3 := member[1:];

	// Membuat slice dari array yang sudah ada (dari pertama ke 3)
	memberSlice4 := member[:3];

	//! Jika menngubah slicenya maka nilai aslinya akan ikut berubah

	// Slice mirip dengan array, tapi tidak perlu jumlah itemnya
	fmt.Println(member);
	fmt.Println(memberSlice);
	fmt.Println(memberSlice2);
	fmt.Println(memberSlice3);
	fmt.Println(memberSlice4);
}