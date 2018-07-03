package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/minodisk/thriftast/parser"
)

func main() {
	file, err := os.Open("test.thrift")
	if err != nil {
		panic(err)
	}

	j, _ := json.MarshalIndent(parser.Parse(file), "", "  ")
	fmt.Printf("%s\n", j)
}
