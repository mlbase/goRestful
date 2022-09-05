package main

import "os"

func main() {

	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}
