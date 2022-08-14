package helper

import "testing"

func TestSayHelloWord(t *testing.T) {
	result := SayHello("Danish")
	if result != "Hello Danish" {
		// unit test failed
		panic("Return is NOT Hello Danish")
	}
}
