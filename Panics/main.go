package main

import (
	"fmt"
	"log"
)

func main() {
	divide(4, 0)
	fmt.Println("after panic")
}

func divide(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic happened!!!! \n", err)
		}
	}()
	fmt.Println(a / b)
}
