package main

import (
	"fmt"
	"os"
	"os/exec"
)

//`go run main.go` run /bin/bash
func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("run args not provided")

	}
}

func run() {
	fmt.Printf("%v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
