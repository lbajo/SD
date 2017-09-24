package main

import (
	"errors"
	"fmt"
	"logiclog"
	"os"
)

func main() {
	args := os.Args[1:]
	num_args := len(args)

	if num_args < 1 {
		fmt.Println(errors.New("Hay que introducir mínimo un fichero"))
	}

	for i := 1; i <= num_args; i++ {
		logiclog.ProcessFiles(os.Args[i])
	}

	logiclog.Order()
}
