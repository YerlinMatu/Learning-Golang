package main;

import "fmt";

func main() {
	var age int;
	var name string;
	fmt.Print(" What is your name? :");
	fmt.Scanf("%s", &name);
	fmt.Print(" Hello "+ name +", What is your age? :");
	fmt.Scanf("%d", &age);
	fmt.Println("Great!, Your age is", age);
	fmt.Println("Maybe you are of ~", (2018 - age));
}