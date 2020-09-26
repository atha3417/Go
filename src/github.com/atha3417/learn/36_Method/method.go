package main;

import ("fmt");

type Member struct {
	name string
	age int
}

// Method
func (m Member) getInfo() {
	fmt.Println("Saya method dari Member");
}

func main() {
	var user = Member{"Atha", 14};
	user.getInfo();
	fmt.Println(user);
}