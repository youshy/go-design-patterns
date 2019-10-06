package main

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Printf("Hello World!\n")
	time.Sleep(1 * time.Second)
}

func timer(a func()) {
	fmt.Printf("Start: %s\n", time.Now())
	a()
	fmt.Printf("End: %s\n", time.Now())
}

func main() {
	fmt.Printf("Type: %T\n", myFunc)
	timer(myFunc)
}
