package main;

import "fmt";

// in loop for (numbers of pairs from 1 to 20).
func main() {
	for i := 0; i <= 20; i++ {
		if i % 2 == 0 && i != 0 {
			fmt.Println(i);
		}
	}
}