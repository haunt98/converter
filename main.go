package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(`go run main.go "This is a sentence"`)
		return
	}

	text := args[1]
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, "'", "_")
	text = strings.ReplaceAll(text, `"`, "_")
	text = strings.ReplaceAll(text, " ", "_")

	fmt.Println(text)
}
