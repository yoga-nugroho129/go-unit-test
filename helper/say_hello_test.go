package helper

import (
	"fmt"
	"runtime"
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

// #2 HANDLE ERROR (FAIL TEST)
func Test2SayHello2(t *testing.T) {
	result := SayHello2("Danish")
	if result != "Hello Danish" {
		t.Fail()
	}

	fmt.Println("Test2SayHello2 DONE") // tetap dieksekusi
}

func Test3SayHello2(t *testing.T) {
	result := SayHello2("Danish")
	if result != "Hello Danish" {
		t.FailNow()
	}

	fmt.Println("Test3SayHello2 DONE") // tidak dieksekusi
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

// #3 ASSERT MODULE
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

// #4 SKIP TEST
func TestSkipSayHello2(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Mac OS")
	}

	result := SayHello2("Danish") // tidak dieksekusi karena skip
	require.Equal(t, "Hello Danish", result)
}

// #5 testing.M as BEFORE - AFTER TEST
func TestMain(m *testing.M) {
	fmt.Println("===Mulai Test===")

	m.Run()

	fmt.Println("===Selesai Test===")
}

// #6 SUBTEST
func TestSubtest(t *testing.T) {
	t.Run("Danish", func(t *testing.T) {
		result := SayHello("Danish")
		assert.Equal(t, "Hello Danish", result)
	})
	t.Run("Cici", func(t *testing.T) {
		result := SayHello("Cici")
		assert.Equal(t, "Hello Cici", result)
	})
}

// #7 TABLE TEST
func TestTableSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Danish",
			request:  "Danish",
			expected: "Hello Danish",
		},
		{
			name:     "Cici",
			request:  "Cici",
			expected: "Hello Cici",
		},
		{
			name:     "Yoga",
			request:  "Yoga",
			expected: "Hello Yoga",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

// #9 BENCHMARK
func BenchmarkSayHelloDanish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Danish")
	}
}

func BenchmarkSayHelloCiciFinansia(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Cici Finansia")
	}
}

// #10 SUB BENCHMARK
func BenchmarkSub(b *testing.B) {
	b.Run("Danish", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Danish")
		}
	})

	b.Run("DanishFull", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("M. Danish Alfaridzi")
		}
	})
}

// #11 TABLE BENCHMARK
func BenchmarkTableSayHello(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloDanish",
			request: "Danish",
		},
		{
			name:    "HelloAlfaridzi",
			request: "Alfaridzi",
		},
		{
			name:    "HelloCici",
			request: "Cici",
		},
		{
			name:    "HelloYoga",
			request: "Yoga",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)
			}
		})
	}
}
