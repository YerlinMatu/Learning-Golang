package main; 

import "fmt";

type User interface {
	Autorization() int;
	Name() string; 
}

type Admin struct {
	name string;
}

type Editor struct {
	name string;
}

func (this Admin) Autorization() int {
	return 4;
}

func (this Admin) Name() string {
	return this.name;
}

func (this Editor) Autorization() int {
	return 2;
}

func (this Editor) Name() string {
	return this.name;
}

func stripe() {
	 fmt.Println("----------------");
}

func auth(profile User) string {
	if (profile.Autorization() == 4 ) {
		return profile.Name() + ", You have permits of administration.";
	} else if (profile.Autorization() <=  2) {
		return profile.Name() +", You have permits of editor.";
	} else {
		return "Without access to the system.";
	}
}

func main() {
	admin := Admin{"Yerlin"};
	editor := Editor{"John Doe"};

	fmt.Println(auth(admin));
	stripe();
	fmt.Println(auth(editor));
}