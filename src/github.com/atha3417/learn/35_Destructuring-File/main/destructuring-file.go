package main;

import (
	"fmt"
	"github.com/atha3417/learn/35_Destructuring-File/models"
);

func main() {
	var user1 = models.Member {Name: "Atha", Age: 14};
	fmt.Println(user1);
}