package main;

import "fmt";

func main() {
	ageOfTiger := 20;
	ageOfBunny := 20;
	
	if ageOfBunny > ageOfTiger {
		fmt.Println("Bunny is higher");
	} else if ageOfTiger > ageOfBunny {
		fmt.Println("Tiger is higher");
	} else {
		fmt.Println("Bunny and Tiger are equal.");
	}
}