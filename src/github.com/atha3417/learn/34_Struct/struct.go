package main;

import ("fmt");

// Struct
type Member struct {
	name string
	age int
}

func main() {
	// var user1 = Member {"Atha", 14};
	var user1 = Member {age: 14, name: "Tsaqif"};
	
	// Value type || Reference type
	var user2 = user1;
	
	user1.name = "Atha Tsaqif";

	fmt.Println(user1);
	fmt.Println(user2);
}