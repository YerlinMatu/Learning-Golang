package main

import "fmt"

func main() {
	arrOfFruits := [4]string{"Grape", "Cherry", "Strawberry", "Raspberry"};
	sliceOfrrys := arrOfFruits[1:4]; // 1:4 (from 1 to 4 position).
	fmt.Println(sliceOfrrys);
}
