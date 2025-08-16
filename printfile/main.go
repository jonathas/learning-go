package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	openFile(os.Args[1])
}

func openFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}