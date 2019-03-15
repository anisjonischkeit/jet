package generator

var SqlBuilderTableTemplate = `package table

import "github.com/sub0Zero/go-sqlbuilder/sqlbuilder"

type {{.ToGoStructName}} struct {
	sqlbuilder.Table
	
	//Columns
{{- range .Columns}}
	{{.ToGoFieldName}} sqlbuilder.NonAliasColumn
{{- end}}

	AllColumns sqlbuilder.ColumnList
}

var {{.ToGoVarName}} = &{{.ToGoStructName}}{
	Table: *sqlbuilder.NewTable("{{.DatabaseInfo.SchemaName}}", "{{.Name}}", {{.ToGoColumnFieldList ", "}}),
	
	//Columns
{{- range .Columns}}
	{{.ToGoFieldName}}: {{.ToGoVarName}},
{{- end}}

	AllColumns: sqlbuilder.ColumnList{ {{.ToGoColumnFieldList ", "}} },
}

var (
{{- range .Columns}}
	{{.ToGoVarName}} = sqlbuilder.IntColumn("{{.Name}}", {{if .IsNullable}}sqlbuilder.Nullable{{else}}sqlbuilder.NotNullable{{end}})
{{- end}}
)
`

var DataModelTemplate = `package model

{{range .GetImports}}
	import "{{.}}"
{{end}}

type {{.ToGoModelStructName}} struct {
{{- range .Columns}}
	{{.ToGoDMFieldName}} {{.ToGoType}} {{if .IsUnique}}` + "`sql:\"unique\"`" + ` {{end}}
{{- end}}
}
`