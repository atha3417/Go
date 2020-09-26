package main;

import ("fmt");

func bayarUtang(uang int, utang int) (pesan string) {
	if uang >= utang {
		pesan = "Utang lunas... tunggu kembaliannya";
	} else if uang == utang {
		pesan = "Uangnya pass... lunas";
	} else {
		pesan = "Uang kamu tidak cukup";
	}

	return pesan;
}

func main() {
	fmt.Println(bayarUtang(1000, 5000));
}