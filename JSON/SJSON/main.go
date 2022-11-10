package main

import (
	"fmt"
	"github.com/tidwall/sjson"
)

func main() {

	json := `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Fear and Loathing in Las Vegas",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

	value, _ := sjson.Set(json, "name.first", "Mikhail")
	fmt.Println(value)
}
