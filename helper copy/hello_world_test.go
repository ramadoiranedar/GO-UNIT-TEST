package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Faill") // testing
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
