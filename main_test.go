package go_docker

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := hello()
	if result != "Hello World!" {
		t.Fatalf("hello() didn't output \"Hello World!\", got:\n%q", result)
	}
}