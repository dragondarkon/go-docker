package main

import "fmt"

func main() {
	fmt.Println(hello())
}

// hello *should return Hello World!"
func hello() string {
	return "Hello World!"
}

func FizzBuzz(number int) string {
	if number == 1 {
		return "1"
	}
	return "2"
}
