package main

import "os"

func main() {
	openFile("Invalid.txt")

	println("Done")
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

}
