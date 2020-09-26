package main;

import ("fmt");

func main() {
	// member := [5] string {"Atha", "Nafisa", "Fathan", "Aiman", "Fayza"};
	// sliceMember := member[:3];

	// Membuat slice dengan make(tipe data, panjang, [kapasitas])
	member := make([]string, 5);

	fmt.Println(member);
	fmt.Println("Length", len(member));
	fmt.Println("Kapasitas", cap(member));
}