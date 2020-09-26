package main;

import ("fmt");

func bayarUtang(uang int, utang int) (pesan string) {
	// Switch Case mengatasi jika terlalu banyak pengujian
	switch {
		case uang > utang:
			pesan = "Utang lunas... tunggu kembaliannya";
		case uang == utang:
			pesan = "Uangnya pass... lunas";
	}

	return pesan;
}

func main() {
	fmt.Println(bayarUtang(1000, 5000));
}