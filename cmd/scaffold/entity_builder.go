package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
	"text/template"

	"github.com/ettle/strcase"
	"github.com/urfave/cli/v2"
)

var mapOptionFields = map[string]string{
	"valueobject.UUID":        "string",
	"valueobject.Name":        "string",
	"valueobject.Email":       "string",
	"valueobject.Cpf":         "string",
	"valueobject.CarPlate":    "string",
	"valueobject.Hash":        "string",
	"valueobject.AccountType": "string",
	"valueobject.Date":        "int64",
}

var mapOptionBuilders = map[string]string{
	"valueobject.UUID":        "valueobject.UUIDFromString(id)",
	"valueobject.Name":        "valueobject.NameFromString(name)",
	"valueobject.Email":       "valueobject.EmailFromString(email)",
	"valueobject.Cpf":         "valueobject.CpfFromString(cpf)",
	"valueobject.CarPlate":    "valueobject.CarPlateFromString(carPlate)",
	"valueobject.Hash":        "valueobject.LoadHashFromString(hash)",
	"valueobject.AccountType": "valueobject.AccountTypeFromString(accountType)",
	"valueobject.Date":        "valueobject.DateFromUnix(date)",
}

const (
	entityBuilderTemplate = "assets/entity_builder.tmpl"
)

var buildFuncTempl = `
func {{.EntityName}}With{{.PCFieldName}}({{.CCParamName}} {{.ParamType}}) {{.CCEntityName}}Option {
	return func(opt *{{.EntityName}}) error {
		var err error
		opt.{{.FieldName}}, err = {{.FieldBuilder}}
		return nil
	}
}

`

type EntityBuilderParams struct {
	EntityName   string
	CCEntityName string
	Funcs        string
}

type EntityBuilderFuncParams struct {
	EntityName   string
	PCFieldName  string
	CCParamName  string
	ParamType    string
	CCEntityName string
	FieldName    string
	FieldBuilder string
}

func createEntityBuilder() *cli.Command {
	return &cli.Command{
		Name:  "entitybuilder",
		Usage: "Create a new entity builder",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Aliases:  []string{"n"},
			},
		},
		Action: func(c *cli.Context) error {
			pwd, _ := os.Getwd()
			buf := bytes.NewBuffer(nil)
			if err := CreateBuildFuncFromEntity(buf, pwd, c.String("name")); err != nil {
				return err
			}
			return createTemplateFile(
				fmt.Sprintf("%s/%s_builder.go", domainPath, toSnakeCase(c.String("name"))),
				template.Must(template.ParseFS(res, entityBuilderTemplate)),
				EntityBuilderParams{
					EntityName:   c.String("name"),
					CCEntityName: strcase.ToCamel(c.String("name")),
					Funcs:        buf.String(),
				},
			)
		},
	}
}

func CreateBuildFuncFromEntity(buf *bytes.Buffer, path string, entityName string) error {
	path = path + "/pkg/domain"
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	for _, pkg := range pkgs {
		checker := &types.Config{IgnoreFuncBodies: true,
			FakeImportC: true,
			Error:       nil,
			Importer:    importer.ForCompiler(fset, "source", nil),
		}
		files := make([]*ast.File, 0)
		for fileName, file := range pkg.Files {
			if strings.HasSuffix(fileName, "_test.go") {
				continue
			}
			files = append(files, file)
		}
		info := &types.Info{
			Defs:  map[*ast.Ident]types.Object{},
			Uses:  map[*ast.Ident]types.Object{},
			Types: map[ast.Expr]types.TypeAndValue{},
		}
		pkg, err := checker.Check(path, fset, files, info)
		if err != nil {
			return err
		}
		obj := pkg.Scope().Lookup("Account")
		if obj == nil {
			return fmt.Errorf("object not found")
		}
		if _, ok := obj.(*types.TypeName); !ok {
			return fmt.Errorf("%v is not a named type", obj)
		}
		iface, ok := obj.Type().Underlying().(*types.Struct)
		if !ok {
			return fmt.Errorf("type %v is a %T, not an interface",
				obj, obj.Type().Underlying())
		}
		numFields := iface.NumFields()
		fieldNames := make([]string, numFields)
		fieldTypes := make([]string, numFields)
		fieldBuilders := make([]string, numFields)
		for i := 0; i < numFields; i++ {
			field := iface.Field(i)
			fieldNames[i] = field.Name()
			fieldTypes[i] = getPackagePath(field.Type().String())
			fieldBuilders[i] = getFieldBuilders(field.Type().String())
		}

		result := getWithFuncName(entityName, fieldNames, fieldTypes, fieldBuilders)
		buf.Write([]byte(result))
	}
	return nil
}

func getPackagePath(pkgName string) string {
	s := strings.Split(pkgName, "/")
	f, ok := mapOptionFields[s[len(s)-1]]
	if !ok {
		panic("package not found")
	}
	return f
}

func getFieldBuilders(pkgName string) string {
	s := strings.Split(pkgName, "/")
	f, ok := mapOptionBuilders[s[len(s)-1]]
	if !ok {
		panic("package not found")
	}
	return f
}

func getWithFuncName(entityName string, fieldNames, fieldTypes, fieldBuilders []string) string {
	tmpl := template.Must(template.New("buildFunc").Parse(buildFuncTempl))
	buf := bytes.NewBuffer(nil)
	for i := 0; i < len(fieldNames); i++ {
		tmpl.Execute(buf, EntityBuilderFuncParams{
			EntityName:   entityName,
			PCFieldName:  strcase.ToGoPascal(fieldNames[i]),
			CCParamName:  strcase.ToCamel(fieldNames[i]),
			ParamType:    fieldTypes[i],
			CCEntityName: strcase.ToCamel(entityName),
			FieldName:    fieldNames[i],
			FieldBuilder: fieldBuilders[i],
		})
	}
	return buf.String()
}
