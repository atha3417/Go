package main;

import ("fmt");

var member = map[int] string {
	34566: "Atha",
	45678: "Tsaqif",
	98767: "Muhtasaq",
};

func main() {
	if checkMember(98767) {
		fmt.Println("Member dengan id ini ada");
	} else {
		fmt.Println("Member dengan id ini tidak ada");
	}

	// Delete -> delete(map, key)
	delete(member, 98767);

	fmt.Println(member);
}

func checkMember(id int) (exist bool) {
	_, exist = member[id];
	return;
}