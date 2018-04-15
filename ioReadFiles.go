package main;

import (
	"io/ioutil"
	"fmt"
);

func main() {
	datafile, err := ioutil.ReadFile("./zfile.txt");
	if err != nil {
		fmt.Println("Exist a err.");
	} 
	fmt.Println(string(datafile));
}