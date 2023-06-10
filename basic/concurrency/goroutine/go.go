package goroutine

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func Run() {
	printSomething("start...")

	// from main routine
	printSomething("Hello from main routine!")

	// from go routine
	go printSomething("Hello from go routine!")

	// to waiting for go routine to finish
	time.Sleep(1 * time.Second)

	printSomething("exit...")
}
