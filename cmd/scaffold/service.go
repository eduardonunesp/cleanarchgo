package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/urfave/cli/v2"
)

var (
	//go:embed assets
	res embed.FS
)

const (
	serviceTemplate     = "assets/service.tmpl"
	serviceTestTemplate = "assets/service_test.tmpl"
)

type ServiceParams struct {
	ServiceName string
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

func createService() *cli.Command {
	return &cli.Command{
		Name:  "new_service",
		Usage: "Create a new service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Aliases:  []string{"n"},
			},
		},
		Action: func(c *cli.Context) error {
			tmpl, err := template.ParseFS(res, serviceTemplate)
			if err != nil {
				return err
			}
			buf := new(bytes.Buffer)
			if err := tmpl.Execute(buf, ServiceParams{
				ServiceName: c.String("name"),
			}); err != nil {
				return err
			}
			if err := createTemplateFile(
				fmt.Sprintf("pkg/service/%s.go", toSnakeCase(c.String("name"))),
				tmpl, ServiceParams{
					ServiceName: c.String("name"),
				},
			); err != nil {
				return err
			}
			return createTemplateFile(
				fmt.Sprintf("pkg/service/%s_test.go", toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, serviceTestTemplate)),
				ServiceParams{
					ServiceName: c.String("name"),
				},
			)
		},
	}
}
