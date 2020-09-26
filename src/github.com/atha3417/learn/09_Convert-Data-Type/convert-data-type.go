package main;

import (
	"fmt"
	"strconv"
);

func main() {
	// Integer to String
	gaji := 1000;
	fmt.Println("Gaji saat ini..." + strconv.Itoa(gaji));

	// String to Integer
	gajiBesok := "1000";
	gajiInt, _ := strconv.Atoi(gajiBesok);
	bonus := gajiInt + 5000;
	fmt.Println("Gaji besok adalah...", bonus);
}