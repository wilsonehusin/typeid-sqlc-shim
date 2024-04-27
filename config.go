package main

import (
	"fmt"
	"strings"
)

type Config struct {
	Package Package `json:"package"`
	Models  []Model `json:"models"`
}

func (c Config) Valid() bool {
	p := c.Package.Valid()

	m := true
	for _, model := range c.Models {
		if !model.Valid() {
			m = false
		}
	}
	return p && m
}

type Package struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (p Package) Valid() bool {
	hasError := false
	if p.Name == "" {
		errorLog("package.name is required\n")
		hasError = true
	}
	if p.Path == "" {
		errorLog("package.path is required\n")
		hasError = true
	}
	return !hasError
}

type Model struct {
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
	Table  string `json:"table"`
}

func (m Model) SQLColumn() string {
	return fmt.Sprintf("%s.id", m.Table)
}

func (m Model) PrefixStr() string {
	if m.Prefix != "" {
		return m.Prefix
	}
	return strings.ToLower(m.Name)
}

func (m Model) TypePrefix() string {
	return fmt.Sprintf("%sPrefix", m.Name)
}

func (m Model) TypeID() string {
	return fmt.Sprintf("%sID", m.Name)
}

func (m Model) Valid() bool {
	hasError := false
	if m.Name == "" {
		errorLog("model.name is required\n")
		hasError = true
	}
	if m.Prefix == "" {
		if p := m.PrefixStr(); p != "" {
			warnLog("model.prefix is empty, defaulting to \"%s\"\n", p)
		}
	}
	return !hasError
}
