package main;

import (
	"fmt"
	"strconv"
);

func getBiography(umur int, nama string, status string) (string, string) {
	return nama + " adalah seorang " + status, "Umurnya sekarang... " + strconv.Itoa(umur);
}

func main() {
	basicInfo, ageInfo := getBiography(21, "Dyballa", "Pemain Bola");
	fmt.Println(basicInfo, ageInfo);
}