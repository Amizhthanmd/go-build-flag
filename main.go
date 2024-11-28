package main

import (
	"fmt"
	"os"
)

var (
	// Metadata variables set during the build process
	version   string
	buildDate string
	commit    string
)

func main() {
	fmt.Printf("Application Metadata:\n")
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Build Date: %s\n", buildDate)
	fmt.Printf("Commit Hash: %s\n", commit)

	fmt.Println("\nRunning application...")
	if err := runApp(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
		os.Exit(1)
	}
}

func runApp() error {
	fmt.Println("Application is up and running!")
	return nil
}
