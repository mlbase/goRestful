package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait()
}
