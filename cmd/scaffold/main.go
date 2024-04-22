package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func toSnakeCase(s string) string {
	var result string

	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' {
			result += "_"
		}

		result += string(v)
	}

	return strings.ToLower(result)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	app := &cli.App{
		Name: "scaffold",
		Commands: []*cli.Command{
			createService(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
