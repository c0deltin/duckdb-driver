package duckdb

import (
	"gorm.io/gorm/migrator"
)

type Migrator struct {
	migrator.Migrator
}

var typeAliasMap = map[string][]string{
	"int8":      {"bigint"},
	"long":      {"bigint"},
	"int4":      {"integer"},
	"int":       {"integer"},
	"signed":    {"integer"},
	"int2":      {"smallint"},
	"short":     {"smallint"},
	"int1":      {"tinyint"},
	"bool":      {"boolean"},
	"logical":   {"boolean"},
	"bytea":     {"blob"},
	"binary":    {"blob"},
	"varbinary": {"blob"},
	"float8":    {"double"},
	"numeric":   {"double"},
	"decimal":   {"double"},
	"float4":    {"real"},
	"float":     {"real"},
	"char":      {"varchar"},
	"bpchar":    {"varchar"},
	"text":      {"varchar"},
	"string":    {"varchar"},
}

func (m Migrator) GetTypeAliases(databaseTypeName string) []string {
	return typeAliasMap[databaseTypeName]
}
