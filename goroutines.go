package main;

import "fmt";
import "strings";

func main() {
	nameSuperSlooow("Yerlinson");
}

func nameSuperSlooow(name string) {
	letters := strings.Split(name, "");
	for _,letter := range(letters) {
		fmt.Println(letter);
	}
}