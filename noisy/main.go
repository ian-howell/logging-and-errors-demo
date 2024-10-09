package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	a, b := 10, 0
	// Simulate 15 failures
	for i := 0; i < 15; i++ {
		_, err := divide(a, b)
		if err != nil {
			log.Printf("could not divide %d by %d: %v", a, b, err)
		}
	}

	// Follow it up with a single success
	b = 2
	result, err := divide(a, b)
	if err != nil {
		log.Printf("could not divide %d by %d: %v", a, b, err)
	}
	log.Printf("result: %d\n", result)
}

func divide(a, b int) (int, error) {
	log.Println("dividing")
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	log.Println("divided")
	return a / b, nil
}
