package main

import (
	"fmt"
	"text/template"

	"github.com/urfave/cli/v2"
)

const (
	repositoryTemplate          = "assets/repository.tmpl"
	repositoryTestTemplate      = "assets/repository_test.tmpl"
	repositoryInterfaceTemplate = "assets/repository_interface.tmpl"
	repositoryPath              = "pkg/infra/repository"
	repositoryTypePath          = "types.go"
)

type repositoryParams struct {
	RepositoryName string
}

func createRepository() *cli.Command {
	return &cli.Command{
		Name:  "repository",
		Usage: "Create a new repository in the infra/repository layer",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Aliases:  []string{"n"},
			},
		},
		Action: func(c *cli.Context) error {
			if err := appendTemplateFile(
				fmt.Sprintf("%s/%s", repositoryPath, repositoryTypePath),
				template.Must(template.ParseFS(res, repositoryInterfaceTemplate)),
				repositoryParams{
					RepositoryName: c.String("name"),
				},
			); err != nil {
				return err
			}
			if err := createTemplateFile(
				fmt.Sprintf("%s/%s.go", repositoryPath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, repositoryTemplate)),
				repositoryParams{
					RepositoryName: c.String("name"),
				},
			); err != nil {
				return err
			}
			return createTemplateFile(
				fmt.Sprintf("%s/%s_test.go", repositoryPath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, repositoryTestTemplate)),
				repositoryParams{
					RepositoryName: c.String("name"),
				},
			)
		},
	}
}
