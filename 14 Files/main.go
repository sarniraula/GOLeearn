package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println(dataByte)         // output => [84 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105 110 32 97 32 102 105 108 101 45 32 71 111 32 76 97 110 103 32 84 117 116 111 114 105 97 108]
	fmt.Println(string(dataByte)) // output => This needs to go in a file- Go Lang Tutorial
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Files in Go lang")

	content := "This needs to go in a file- Go Lang Tutorial"

	file, err := os.Create("./myGoFIle.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./myGoFIle.txt")
}
