package concurrent

import (
	"fmt"
	"runtime"
	"testing"
)

func sayHello() {
	fmt.Println("Hello, World!")
}

func TestConcurrentFunction(t *testing.T) {
	go sayHello()
	runtime.Gosched()
	fmt.Println("Main function completed.")
}
