package main;

import "fmt";

type Person struct {
	name string;
}

type Developer struct {
	Person;
	language string;
	web, desktop, mobile bool;
}

func (this *Developer) changeLanguage(newLang string) {
	this.language = newLang;
}

func main() {
	matu := Developer{Person{"Yerlin"}, "JavaScript", true, true, true};
	matu.changeLanguage("Golang");
	fmt.Println(matu);
}
