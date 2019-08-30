package go_docker

import "fmt"

func main() {
	fmt.Println(hello())
}

// hello *should return Hello World!"
func hello() string {
	return "Hello World!"
}