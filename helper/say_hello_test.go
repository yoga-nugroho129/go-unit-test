package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestAsert1SayHello2(t *testing.T) {
	result := SayHello2("Danish")
	assert.Equal(t, "Hello Danish", result, "TEST ASSERT.EQUALDONE") // ASERT jika gagal lanjut memanggil Fail()
	fmt.Println("===ASSERT.EQUAL SELESAI===")                        // tetap dieksekusi
}

func TestAsert2SayHello2(t *testing.T) {
	result := SayHello2("Danish")
	require.Equal(t, "Hello Danish", result, "TEST ASSERT.EQUALDONE") // REQUIRE jika gagal lanjut memanggil FailNow()
	fmt.Println("===REQUIRE.EQUAL SELESAI===")                        // tidak dieksekusi
}
