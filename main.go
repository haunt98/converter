package main

import (
	"fmt"
	"os"

	"github.com/haunt98/converter/pkg/converter"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(`go run main.go "Your string goes here"`)
		return
	}

	fmt.Println(converter.Convert(args[1]))
}
