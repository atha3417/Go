package main;

import ("fmt");

func main() {
	// Mengeluarkan isi array dengan for loop
	members := [3] string {"Atha", "Fathan", "Nafisa"};
	for i := 0; i < len(members); i++ {
		fmt.Println("Member ke", i, "adalah " + members[0]);
	}
}