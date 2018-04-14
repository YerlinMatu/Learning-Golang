package main;

import "fmt";
import "strconv"

func main() {
	age := 20;
	birthday := "27";

	fmt.Println("I am writing with Golang with "+ strconv.Itoa(age) +" years of age.");
	fmt.Println(strconv.Atoi(birthday));
	/*
		strconv.Itoa(variable) : Transform "string" to integer.
		strconv.Atoi(variable) : Transform integer a "string". 
	*/	
}