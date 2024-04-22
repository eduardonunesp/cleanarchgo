package main

import (
	"fmt"
	"text/template"

	"github.com/urfave/cli/v2"
)

const (
	gatewayTemplate = "assets/gateway.tmpl"
	gatewayPath     = "pkg/infra/gateway"
)

type GatewayParams struct {
	GatewayName string
}

func createGateway() *cli.Command {
	return &cli.Command{
		Name:  "gateway",
		Usage: "Create a new gateway at pkg/infra/gateway",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Aliases:  []string{"n"},
			},
		},
		Action: func(c *cli.Context) error {
			return createTemplateFile(
				fmt.Sprintf("%s/%s.go", gatewayPath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, gatewayTemplate)),
				GatewayParams{
					GatewayName: c.String("name"),
				},
			)
		},
	}
}
