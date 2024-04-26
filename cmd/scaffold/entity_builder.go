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

var mapWithFuncParam = map[string]string{
	"valueobject.UUID":        "string",
	"valueobject.Name":        "string",
	"valueobject.Email":       "string",
	"valueobject.Cpf":         "string",
	"valueobject.CarPlate":    "string",
	"valueobject.Hash":        "string",
	"valueobject.AccountType": "string",
	"valueobject.Date":        "int64",
	"valueobject.Coord":       "string",
}

var mapWithFuncFieldBuilder = map[string]string{
	"valueobject.UUID":        "valueobject.UUIDFromString",
	"valueobject.Name":        "valueobject.NameFromString",
	"valueobject.Email":       "valueobject.EmailFromString",
	"valueobject.Cpf":         "valueobject.CpfFromString",
	"valueobject.CarPlate":    "valueobject.CarPlateFromString",
	"valueobject.Hash":        "valueobject.LoadHashFromString",
	"valueobject.AccountType": "valueobject.AccountTypeFromString",
	"valueobject.Date":        "valueobject.DateFromUnix",
	"valueobject.Coord":       "valueobject.BuildCoord",
}

const (
	entityBuilderTemplate = "assets/entity_builder.tmpl"
)

var buildGeneralFuncTempl = `
func {{.EntityName}}With{{.PCFieldName}}({{.CCParamName}} {{.ParamType}}) {{.CCEntityName}}Option {
	return func(opt *{{.EntityName}}) error {
		var err error
		opt.{{.FieldName}}, err = {{.FieldBuilder}}({{.CCParamName}})
		return err
	}
}
`

var buildIDFuncTempl = `
func {{.EntityName}}WithNew{{.PCFieldName}}() {{.CCEntityName}}Option {
	return func(opt *{{.EntityName}}) error {
		opt.{{.FieldName}} = valueobject.MustUUID()	
		return nil
	}
}
`

var buildRawHashFuncTempl = `
func {{.EntityName}}WithEncoded{{.PCFieldName}}({{.CCParamName}} string) {{.CCEntityName}}Option {
	return func(opt *{{.EntityName}}) error {
		var err error
		opt.hash, err = valueobject.BuildHashFromString({{.CCParamName}}, nil)
		return err
	}
}
`

var buildCoordFuncTempl = `
func {{.EntityName}}With{{.PCFieldName}}(x, y string) {{.CCEntityName}}Option {
	return func(p *{{.EntityName}}) error {
		var err error
		p.{{.FieldName}}, err = valueobject.BuildCoord(x, y)
		return err
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
		obj := pkg.Scope().Lookup(entityName)
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
		for i := 0; i < numFields; i++ {
			field := iface.Field(i)
			fieldNames[i] = field.Name()
			fieldTypes[i] = field.Type().String()
		}

		result := getWithFuncName(entityName, fieldNames, fieldTypes)
		buf.Write([]byte(result))
	}
	return nil
}

func getPackageName(path string) string {
	s := strings.Split(path, "/")
	return s[len(s)-1]
}

func getWithFuncParam(pkgName string) string {
	s := getPackageName(pkgName)
	f, ok := mapWithFuncParam[s]
	if !ok {
		panic(fmt.Sprintf("package not found func param %s", s))
	}
	return f
}

func getWithFuncFieldBuilder(pkgName string) string {
	s := getPackageName(pkgName)
	f, ok := mapWithFuncFieldBuilder[s]
	if !ok {
		panic(fmt.Sprintf("package not found at field builder %s", s))
	}
	return f
}

func getWithFuncName(entityName string, fieldNames, fieldTypes []string) string {
	generalWithFuncTmpl := template.Must(template.New("generalBuildFunc").Parse(buildGeneralFuncTempl))
	IdWithFuncTmpl := template.Must(template.New("iDBuildFunc").Parse(buildIDFuncTempl))
	EncodeHashWithFuncTmpl := template.Must(template.New("encodeHashBuildFunc").Parse(buildRawHashFuncTempl))
	CoordWithFuncTmpl := template.Must(template.New("coordBuildFunc").Parse(buildCoordFuncTempl))
	buf := bytes.NewBuffer(nil)
	for i := 0; i < len(fieldNames); i++ {
		shortPackageName := getPackageName(fieldTypes[i])
		entityBuilderParams := EntityBuilderFuncParams{
			EntityName:   entityName,
			PCFieldName:  strcase.ToGoPascal(fieldNames[i]),
			CCParamName:  strcase.ToCamel(fieldNames[i]),
			ParamType:    getWithFuncParam(fieldTypes[i]),
			CCEntityName: strcase.ToCamel(entityName),
			FieldName:    fieldNames[i],
			FieldBuilder: getWithFuncFieldBuilder(fieldTypes[i]),
		}
		var err error
		if shortPackageName != "valueobject.Coord" {
			err = generalWithFuncTmpl.Execute(buf, entityBuilderParams)
		}
		switch shortPackageName {
		case "valueobject.UUID":
			err = IdWithFuncTmpl.Execute(buf, entityBuilderParams)
		case "valueobject.Hash":
			err = EncodeHashWithFuncTmpl.Execute(buf, entityBuilderParams)
		case "valueobject.Coord":
			err = CoordWithFuncTmpl.Execute(buf, entityBuilderParams)
		}
		if err != nil {
			panic(err)
		}
	}
	return buf.String()
}
