package helper

import (
	"fmt"
	"testing"
)

// #1 INTRO
func TestSayHello(t *testing.T) {
	result := SayHello("Danish")
	if result != "Hello Danish" {
		// unit test failed
		panic("Return is NOT Hello Danish")
	}
}

// #2 Handle ERROR
func Test2SayHello2(t *testing.T) {
	result := SayHello2("Danish")
	if result != "Hello Danish" {
		t.Fail()
	}

	fmt.Println("Test2SayHello2 DONE")
}

func Test3SayHello2(t *testing.T) {
	result := SayHello2("Danish")
	if result != "Hello Danish" {
		t.FailNow()
	}

	fmt.Println("Test3SayHello2 DONE")
}

func Test4SayHello2(t *testing.T) {
	result := SayHello2("Danish")
	if result != "Hello Danish" {
		t.Error()
	}

	fmt.Println("Test4SayHello2 DONE")
}

func Test5SayHello2(t *testing.T) {
	result := SayHello2("Danish")
	if result != "Hello Danish" {
		t.Fatal()
	}

	fmt.Println("Test5SayHello2 DONE")
}
