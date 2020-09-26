package main;

import ("fmt");

type Member struct {
	name string
	age int
}

func (m *Member) changeData(newName string, newAge int) {
	m.name = newName;
	m.age = newAge;
}

func main() {
	var user = Member{"Atha", 14};
	user.changeData("Nafisa", 15);
	fmt.Println(user);
}