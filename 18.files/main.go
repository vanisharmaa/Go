package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This is supposed to be in a file"
	file, err := os.Create("./my-go-file.txt")

	checkNilErr(err)

	fmt.Println("file: ", file)   // &{0x1400010e180}
	fmt.Println("file2: ", *file) // {0x1400010e180}

	len, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is: ", len)

	defer file.Close()

	readFile("./my-go-file.txt")
}

func readFile(fileName string) {
	smth, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("This is what is returned: ", smth) // [84 104 105 115 32 105 115 32 115 117 112 112 111 115 101 100 32 116 111 32 98 101 32 105 110 32 97 32 102 105 108 101]
	fmt.Println("As a string: ", string(smth))      // This is supposed to be in a file
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
