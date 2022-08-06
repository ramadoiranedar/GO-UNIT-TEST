package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"  // throw error msg and didn't stop program
	"github.com/stretchr/testify/require" // throw error msg and stop program
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("World")                                          // testing
	assert.Equal(t, "Hello World", result, "Result must be 'Hello World'") // throw error msg and didn't stop program
	fmt.Println("TestHelloWorldAssert with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("LOREM IPSUM DOLORR")                              // testing
	require.Equal(t, "Hello World", result, "Result must be 'Hello World'") // throw error msg and stop program
	fmt.Println("TestHelloWorldRequire with Require Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World") // testing
	if result != "Hello World" {  // required
		// error
		// panic("Result is not 'Hello World'") // NOT RECOMMENDED
		// t.Fail() // didn't stop program
		t.Error("Result must be 'Hello World'") // didn't stop program // with custom msg
	}

	fmt.Println("TestHelloWorld DONE")
}

func TestHelloWorldDendi(t *testing.T) {
	result := HelloWorld("Dendi") // testing
	if result != "Hello Dendi" {  // required
		// error
		// panic("Result is not 'Hello Dendi'") // NOT RECOMMENDED
		// t.FailNow() //  stop program
		t.Fatal("Result must be 'Hello Dendi'") // stop program // with custom msg

	}

	fmt.Println("TestHelloWorldDendi DONE")
}
