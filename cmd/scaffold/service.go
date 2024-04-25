package main

import (
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
	NoReturn    bool
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
			&cli.BoolFlag{
				Name:    "noreturn",
				Aliases: []string{"r"},
				Value:   false,
			},
		},
		Action: func(c *cli.Context) error {
			params := ServiceParams{
				ServiceName: c.String("name"),
				NoReturn:    c.Bool("noreturn"),
			}
			if err := createTemplateFile(
				fmt.Sprintf("%s/%s.go", servicePath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, serviceTemplate)), params,
			); err != nil {
				return err
			}
			return createTemplateFile(
				fmt.Sprintf("%s/%s_test.go", servicePath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, serviceTestTemplate)), params,
			)
		},
	}
}
