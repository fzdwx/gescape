package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	text, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("Error reading stdin: %s\n", err))
		return
	}

	pairs := strings.Split(string(text), "\\n")
	for _, s := range pairs {
		fmt.Println(strings.Replace(s, "\\t", "    ", -1))
	}
}
