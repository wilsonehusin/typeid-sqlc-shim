package main

type SqlcOverride struct {
	Column string `yaml:"column"`
	GoType SqlcGoType `yaml:"go_type"`
}

type SqlcGoType struct {
	Type string `yaml:"type"`
}
