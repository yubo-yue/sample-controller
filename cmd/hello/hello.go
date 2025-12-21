package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	// Create a new FlagSet named "hello"
	fs := pflag.NewFlagSet("hello", pflag.ContinueOnError)

	// Define flags on the FlagSet
	name := fs.String("name", "World", "Name to greet")
	age := fs.Int("age", 0, "Age of the person")
	verbose := fs.Bool("v", false, "Enable verbose output")

	// Parse command-line arguments (excluding the program name)
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Use the flags
	if *verbose {
		fmt.Println("Flag values:")
		fmt.Printf("  name: %s\n", *name)
		fmt.Printf("  age: %d\n", *age)
		fmt.Println()
	}

	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
