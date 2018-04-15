package main;

import "fmt";
import "strings";
import "time";

func main() {
	go nameSuperSlooow("Yerlinson");
	fmt.Println("It's boring to wait!");

	var wait string;
	fmt.Scanln(&wait);
}

func nameSuperSlooow(name string) {
	letters := strings.Split(name, "");
	for _,letter := range(letters) {
		time.Sleep(1000 * time.Millisecond);
		fmt.Println(letter);
	}
}