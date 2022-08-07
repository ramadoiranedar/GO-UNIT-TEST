package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"  // throw error msg and didn't stop program
	"github.com/stretchr/testify/require" // throw error msg and stop program
)

// TABLE BENCHMARK
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "World",
			request: "Wolrd",
		},
		{
			name:    "Damar",
			request: "Damar",
		},
		{
			name:    "Raden Ario Damar",
			request: "Raden Ario Damar",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			Hello(benchmark.request)
		})
	}
}

// SUB BENCHMARK
func BenchmarkSub(b *testing.B) {
	b.Run("World", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("World")
		}
	})
	b.Run("Damar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Damar")
		}
	})
}

// BENCHMARK
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("World")
	}
}

func BenchmarkHelloDamar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Damar")
	}
}

// TABLE TEST
func TestTableHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Raden",
			request:  "Raden",
			expected: "Hello Raden",
		},
		{
			name:     "Damar",
			request:  "Damar",
			expected: "Hello Damar",
		},
		{
			name:     "World",
			request:  "World",
			expected: "Hello World",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// SUB TEST
func TestSubTest(t *testing.T) {
	t.Run("Raden", func(t *testing.T) {
		result := Hello("Raden")
		require.Equal(t, "Hello Raden", result, "Result must be 'Hello Raden'")
	})
	t.Run("Ario", func(t *testing.T) {
		result := Hello("Ario")
		require.Equal(t, "Hello Ario", result, "Result must be 'Hello Ario'")
	})
}

// BEFORE AND AFTER
func TestMain(m *testing.M) {
	fmt.Println("Do stuff before Unit Test, in This PACKAGE 'helper'")
	m.Run() // execute
	fmt.Println("Do Unit Test Now, in This PACKAGE 'helper'")
}

// TEST SKIP
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("This OS is not running on 'windows'")
	}

	result := Hello("Damar")
	require.Equal(t, "Hello Damar", result, "Result must be 'Hello World'") // throw error msg and stop program
}

// ASSERT
func TestHelloAssert(t *testing.T) {
	result := Hello("World")                                               // testing
	assert.Equal(t, "Hello World", result, "Result must be 'Hello World'") // throw error msg and didn't stop program
	fmt.Println("TestHelloAssert with Assert Done")
}

// REQUIRE
func TestHelloRequire(t *testing.T) {
	result := Hello("World")                                                // testing
	require.Equal(t, "Hello World", result, "Result must be 'Hello World'") // throw error msg and stop program
	fmt.Println("TestHelloRequire with Require Done")
}

// ERROR
func TestHello(t *testing.T) {
	result := Hello("World")     // testing
	if result != "Hello World" { // required
		// error
		// panic("Result is not 'Hello World'") // NOT RECOMMENDED
		// t.Fail() // didn't stop program
		t.Error("Result must be 'Hello World'") // didn't stop program // with custom msg
	}

	fmt.Println("TestHello DONE")
}

// FATAL
func TestHelloDendi(t *testing.T) {
	result := Hello("Dendi")     // testing
	if result != "Hello Dendi" { // required
		// error
		// panic("Result is not 'Hello Dendi'") // NOT RECOMMENDED
		// t.FailNow() //  stop program
		t.Fatal("Result must be 'Hello Dendi'") // stop program // with custom msg
	}

	fmt.Println("TestHelloDendi DONE")
}
