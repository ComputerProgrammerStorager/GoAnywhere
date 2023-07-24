package main

import (
	"log"
	"time"
)

// defer a function that returns a function that is called upon exit of the containing function
func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// do works
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("Enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
