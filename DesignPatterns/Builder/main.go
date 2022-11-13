package main

import (
	"Go_jam/DesignPatterns/Builder/computer"
	"fmt"
)

func main() {
	officeCompBuilder := computer.NewOfficeComputerBuilder()
	officeCompBuilder = officeCompBuilder.RAM(32)
	director := computer.NewDirector(officeCompBuilder)
	compB := director.BuildComputer()
	fmt.Println(compB)
}
