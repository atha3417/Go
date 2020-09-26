package main;

import ("fmt");

func main() {
	hewan := [3][2] string {
		{"Macan", "Singa"},
		{"Paus", "Hiu"},
		{"Elang", "Gagak"},
	}

	// mengeluarkan isi array multi dimensi
	for i := 0; i < len(hewan); i++ {
		for j := 0; j < len(hewan[i]); j++ {
			fmt.Println("Hewan " + hewan[i][j]);
		}
		fmt.Println("------------------");
	}
}