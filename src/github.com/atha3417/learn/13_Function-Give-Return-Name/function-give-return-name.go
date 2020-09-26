package main;

import ("fmt");

func getBiography(umur int, nama string, status string) (bio string, infoUmurDalam10 int) {
	infoUmurDalam10 = umur + 10;
	bio = nama + " adalah seorang " + status;
	return;
}

func main() {
	infoDasar, infoUmur := getBiography(21, "Dyballa", "Pemain Bola");
	fmt.Println(infoDasar, ", dan 10 tahun lagi umurnya...", infoUmur);
}