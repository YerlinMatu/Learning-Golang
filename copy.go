package main;

import "fmt";

func main() {
	destination := []int {1, 3, 4, 5};
	// @params remember first 4 is the numbers positions readys. 
	copySource := make([]int, len(destination));
	// Watch data old. 
	fmt.Println("Copy old :", copySource);
	fmt.Println("Destination:", destination);
	// Watch data new. 
	copy(copySource, destination);
	fmt.Println("Copy new :", copySource);
}