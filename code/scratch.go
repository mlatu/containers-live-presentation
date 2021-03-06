package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		panic("use `run` command [arg1 arg2 ...]")
	}

	switch os.Args[1] {
	case "run":
		parent()
	default:
		panic("wat should I do")
	}
}

func parent() {
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func shellOut(command string) {
	parts := strings.Split(command, " ")
	err := exec.Command(parts[0], parts[1:]...).Run()
	if err != nil {
		fmt.Println(err)
	}
}
