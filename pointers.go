package main;

import "fmt";

func main() {
	// is a address inside memory.
	var nick string = "Mr.Matu";
	var pointerOne, pointerTwo *string;

	pointerOne = &nick;
	pointerTwo = &nick;

	fmt.Println("Pointer n#1", *pointerOne);
	fmt.Println("Pointer n#2", *pointerTwo);
	
	fmt.Println("---------------------");
	// Change address of memory.  
	*pointerOne = "Yerlin";

	fmt.Println("Pointer new n#1", *pointerOne);
	fmt.Println("Pointer new n#2", *pointerTwo);
	
	fmt.Println("New nick", nick); 
}