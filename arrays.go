package main;

import "fmt";

func main() {
	numbers := [5]int {2, 4, 6, 8, 10};

	fmt.Println("Numbers of pairs", numbers);
	fmt.Println("Size of array ", len(numbers));
	fmt.Println("First position ", numbers[0]);

	// add +1 to each position of array. 

	for i :=0; i < len(numbers); i++  {
		numbers[i] += 1;
	}

	fmt.Println("Finish loop and numbers not pairs :", numbers);
}