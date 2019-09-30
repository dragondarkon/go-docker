package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := hello()
	if result != "Hello World!" {
		t.Fatalf("hello() didn't output \"Hello World!\", got:\n%q", result)
	}
}

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := FizzBuzz(1) // --> (4)
	if "1" != v {    // --> 5
		t.Error("fizzbuzz of 1 should be '1' but have", v) //
	}
}

func TestInput2ShouldBeDisplay2(t *testing.T) {
	v := FizzBuzz(2)
	if "2" != v {
		t.Error("fizzbuzz of 2 should be '2' but have", v)
	}
}
