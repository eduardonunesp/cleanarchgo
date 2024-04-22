package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

var (
	//go:embed assets
	res embed.FS
)

func main() {
	app := &cli.App{
		Name: "scaffold",
		Commands: []*cli.Command{
			{
				Name:        "create",
				Description: "Create a new service, entity, or repository",
				Subcommands: []*cli.Command{
					createService(),
					createEntity(),
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func createTemplateFile(filename string, tmpl *template.Template, params interface{}) error {
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, params); err != nil {
		return err
	}
	if fileExists(filename) {
		return fmt.Errorf("file %s already exists", filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = file.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

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
