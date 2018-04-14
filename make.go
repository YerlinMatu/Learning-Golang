package	main;

import "fmt";

func main() {
	positionsInitial := 0;
	capacityOfSize := 5;

	slice := make([]int, positionsInitial, capacityOfSize);
	fmt.Println(slice);

	slice = append(slice, 2);
	slice = append(slice, 5);
	slice = append(slice, 100);

	fmt.Println("Push a new elements", slice);
}