package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("this is file handling")
	/////Writing to a file
	content := "The following things musht go inside a file----- pablo is a nice guy"

	file, err := os.Create("./sample.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The length of the file is %v\n Now reading from file \n", length)

	reader("./sample.txt")

}

func reader(file string) {

	databyte, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	fmt.Println("File is found \n the content of file is .......\n", string(databyte))

}
