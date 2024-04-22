package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/urfave/cli/v2"
)

const (
	serviceTemplate     = "assets/service.tmpl"
	serviceTestTemplate = "assets/service_test.tmpl"
	servicePath         = "pkg/service"
)

type ServiceParams struct {
	ServiceName string
}

func createService() *cli.Command {
	return &cli.Command{
		Name:  "service",
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
				fmt.Sprintf("%s/%s.go", servicePath, toSnakeCase(c.String("name"))),
				tmpl, ServiceParams{
					ServiceName: c.String("name"),
				},
			); err != nil {
				return err
			}
			return createTemplateFile(
				fmt.Sprintf("%s/%s_test.go", servicePath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, serviceTestTemplate)),
				ServiceParams{
					ServiceName: c.String("name"),
				},
			)
		},
	}
}
