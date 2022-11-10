package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
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

	if !gjson.Valid(json) {
		panic("JSON IS NOT VALID")
	}

	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		} else if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})

	value := gjson.Get(json, "friends")
	fmt.Println(value.String())

	result, ok := gjson.Parse(json).Value().(map[string]interface{})
	if !ok {
		panic("NOT OKAY PARSING TO MAP")
	}
	fmt.Println(result)
}
