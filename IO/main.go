package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	appendFile()
}

func appendFile() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(" Mikhail"); err != nil {
		panic(err)
	}
}

func writeToFile() {
	data := []byte("my name is Mikhail")
	err := ioutil.WriteFile("test.txt", data, 0777)
	if err != nil {
		fmt.Println(err)
	}

}

func readFile() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
