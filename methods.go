package main;

import "fmt";

type Person struct {
	firstname string;
	lastname string;
	skill string;
	age int;
}
// This the method.
// @param the first this === foo, user, etc.  
func (this Person) getTotalName() string {
	return this.firstname + " " + this.lastname;
}

func (this *Person) setTotalName(firstname string, lastname string) {
	this.firstname = firstname;
	this.lastname = lastname;
}

func stripe() {
	fmt.Println("----------------------------------");
}

func main() {
	yerlin := Person { firstname: "Yerlin", lastname: "Matu", skill: "Software Developer", age: 20}
	// also declare in this way.
	john := Person {"John", "Doe", "Baker", 32 }
	
	fmt.Println(yerlin);
	fmt.Println(john);
	stripe();

	fmt.Println("Profile with details - Yerlin Matu");
	fmt.Println("Name:", yerlin.getTotalName());
	fmt.Println("Skill:", yerlin.skill);
	fmt.Println("Age:", yerlin.age);
	stripe();

	fmt.Println("Profile with details - John Doe");
	john.setTotalName("Mark", "Houses"); // Here change of name.
	fmt.Println("Name:", john.getTotalName());
	fmt.Println("Skill:", john.skill);
	fmt.Println("Age:", john.age);
	stripe();
}