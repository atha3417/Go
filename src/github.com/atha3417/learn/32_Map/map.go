package main;

import ("fmt");

func main() {
	// Map -> key: value
	// Membuat map; map[tipedataKey] tipedataValue {}
	member := map[int] string {
		5678776: "Atha",
		6787657: "Tsaqif",
	};

	member[5678776] = "Fayza";

	// Reference Type
	var newMember = member;

	newMember[5678776] = "Marchelle";

	fmt.Println(member);
	fmt.Println(newMember);
	// fmt.Println(member[5678776]);

	// Mengeluarkan data map dengan for loop
	for id, name := range member {
		fmt.Println(id, ": " + name);
	}
}