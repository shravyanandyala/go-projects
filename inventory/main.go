package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input file provided.")
		os.Exit(1)
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", os.Args[1])
		panic(err)
	}
	fmt.Println(data)
}
