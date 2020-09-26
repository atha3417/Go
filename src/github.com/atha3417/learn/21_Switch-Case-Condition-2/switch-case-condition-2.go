package main;

import ("fmt");

func cekLampu(warna string) (pesan string) {
	// Default artinya else dalam if else
	switch warna {
		case "merah":
			pesan = "Berhenti...";
		case "kuning":
			pesan = "Hati-Hati...";
		case "hijau":
			pesan = "Jalan...";
		default:
			pesan = "Lampu sedang error..."
	}

	return pesan;
}

func main() {
	fmt.Println(cekLampu("hitam"));
}