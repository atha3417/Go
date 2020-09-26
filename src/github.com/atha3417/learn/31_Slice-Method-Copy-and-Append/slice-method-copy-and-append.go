package main;

import ("fmt");

func main() {
	member := [] string {"Atha", "Nafisa", "Fathan", "Aiman", "Fayza"};

	newMember := make([]string, 5);

	// Copy mengkopi slice; copy(dest, src)
	copy(newMember, member);

	// Menambahkan data baru ke dalam slice
	newMember = append(newMember, "Marchelle", "Zahwa");

	newMember[0] = "Muhtasaq";

	fmt.Println(member);
	fmt.Println(newMember);
}