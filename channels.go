package main;

import "fmt";

func main() {
	channel := make(chan string);
	fmt.Println("Initialized!");

	go func(chan string) {
	  for {
		 var name string;
		 fmt.Scanln(&name);
		 channel <- name; // send data to channel.
	  }
	}(channel)
	msg := <- channel; // receive data from channel.
	fmt.Println("Data receive", msg); 
}
