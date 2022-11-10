package main

import (
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("top")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()

	//c := exec.Command("ssh", "-t", "kart", "top")
	//c.Stdin = os.Stdin
	//c.Stdout = os.Stdout
	//c.Stderr = os.Stderr
	//c.Run()
}
