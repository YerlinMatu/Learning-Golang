package main;

import "fmt";

type Person struct {
	name string;
	lastname string;
	skill string;
	age int;
}

func stripe() {
	fmt.Println("----------------------------------");
}

func main() {
	yerlin := Person { name: "Yerlin", lastname: "Matu", skill: "Software Developer", age: 20}
	// also declare in this way.
	john := Person {"John", "Doe", "Baker", 32 }
	
	fmt.Println(yerlin);
	fmt.Println(john);
	stripe();

	fmt.Println("Profile with details - Yerlin Matu");
	fmt.Println("Name:", yerlin.name + yerlin.lastname);
	fmt.Println("Skill:", yerlin.skill);
	fmt.Println("Age:", yerlin.age);
	stripe();

	fmt.Println("Profile with details - John Doe");
	fmt.Println("Name:", john.name + john.lastname);
	fmt.Println("Skill:", john.skill);
	fmt.Println("Age:", john.age);
	stripe();
}