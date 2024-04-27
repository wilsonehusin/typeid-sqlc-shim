package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		errorLog("failed to generate: %v\n", err)
		os.Exit(1)
	}
}

var (
	configPath = flag.String("config", "typeid-sqlc-shim.yaml", "path to the config file")
)

func run() error {
	var conf Config

	raw, err := os.ReadFile(*configPath)
	if err != nil {
		return fmt.Errorf("reading '%s': %w", *configPath, err)
	}
	if err := yaml.Unmarshal(raw, &conf); err != nil {
		return fmt.Errorf("parsing '%s': %w", *configPath, err)
	}

	if err := Render(conf); err != nil {
		return fmt.Errorf("rendering: %w", err)
	}

	infoLog("%s has been written\n", conf.Package.Path)

	s := []SqlcOverride{}
	for _, model := range conf.Models {
		s = append(s, SqlcOverride{
			Column: model.SQLColumn(),
			GoType: SqlcGoType{Type: model.TypeID()},
		})
	}
	m := map[string][]SqlcOverride{
		"overrides": s,
	}
	b, err := yaml.Marshal(m)
	if err != nil {
		return fmt.Errorf("marshalling sqlc.yaml: %w", err)
	}
	infoLog("Append the following to your sqlc.yaml:\n\n%s\n", string(b))
	infoLog("See https://docs.sqlc.dev/en/latest/howto/overrides.html#the-go-type-map for more information.\n")

	return nil
}
