package main

import (
	"fmt"
	"text/template"

	"github.com/urfave/cli/v2"
)

const (
	entityTemplate     = "assets/entity.tmpl"
	entityTestTemplate = "assets/entity_test.tmpl"
	domainPath         = "pkg/domain"
)

type entityParams struct {
	EntityName string
}

func createEntity() *cli.Command {
	return &cli.Command{
		Name:  "entity",
		Usage: "Create a new entity in the domain layer",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Aliases:  []string{"n"},
			},
		},
		Action: func(c *cli.Context) error {
			if err := createTemplateFile(
				fmt.Sprintf("%s/%s.go", domainPath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, entityTemplate)),
				entityParams{
					EntityName: c.String("name"),
				},
			); err != nil {
				return err
			}
			return createTemplateFile(
				fmt.Sprintf("%s/%s_test.go", domainPath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, entityTestTemplate)),
				entityParams{
					EntityName: c.String("name"),
				},
			)
		},
	}
}
