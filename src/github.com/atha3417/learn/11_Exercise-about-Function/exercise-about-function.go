package main;

import (
	"fmt"
	"strconv"
);

func getBiography(umur int, nama string, status string) string {
	return nama + " adalah seorang " + status + " yang saat ini berumur " + strconv.Itoa(umur) + "tahun..";
}

func main() {
	fmt.Println(getBiography(21, "Dyballa", "Pemain Bola"));
	fmt.Println(getBiography(60, "Bill Gates", "Pembuat Microsoft"));
}