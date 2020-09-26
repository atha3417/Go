package main;

import ("fmt");

func changePoint(point *int) {
	// point = 200; // tidak berubah karena //! beda scope
	*point = 200;
}

func main() {
	// Pointer adalah cara menunjuk alamat dari memori variable yang disimpan
	var point = 100;
	var pointAddress *int = &point;

	// *pointAddress = 200; // mengganti nilai yang ada di alamat ini

	changePoint(&point);

	fmt.Println(point);
	fmt.Println(pointAddress);
	// fmt.Println(&point); // mengambil alamat dari point
	// fmt.Println(*pointAddress); // mengambil nilai dari pointAddress
}