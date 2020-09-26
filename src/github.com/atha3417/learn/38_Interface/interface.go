package main;

import (
	"fmt"
	"math"
);

// Interface
type Bentuk interface {
	keliling() float64
	luas() float64
}

type kotak struct {
	panjang, lebar float64
}

type lingkaran struct {
	radius float64
}

func main() {
	var kotak = kotak{5, 10};
	hitungBangunan(kotak);

	var lingkaran = lingkaran{10};
	hitungBangunan(lingkaran);
}

func hitungBangunan(b Bentuk) {
	fmt.Println("Kelilingya..", b.keliling());
	fmt.Println("Kelilingya..", b.luas());
}

func (k kotak) keliling() float64 {
	return (2 * k.panjang) + (2 * k.lebar);
}

func (k kotak) luas() float64 {
	return k.panjang * k.lebar;
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.radius;
}

func (l lingkaran) luas() float64 {
	return math.Pi * l.radius * l.radius;
}