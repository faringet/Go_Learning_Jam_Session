package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string
	Age       int
	IsBlocked bool
}

func main() {
	sv := User{
		"Mikhail",
		28,
		false,
	}
	boolVar, _ := json.Marshal(sv)
	fmt.Println(string(boolVar))
	//json.Unmarshal()
}
