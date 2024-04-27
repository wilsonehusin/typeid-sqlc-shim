package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed render.tmpl
var renderTmpl string

var t = template.Must(template.New("model").Parse(renderTmpl))

func Render(conf Config) error {
	if !conf.Valid() {
		return errors.New("config is invalid")
	}

	f, err := os.Stat(conf.Package.Path)
	if !errors.Is(err, os.ErrNotExist) {
		if f.IsDir() {
			return fmt.Errorf("found '%s' to be a directory", conf.Package.Path)
		}
	}

	d := filepath.Dir(conf.Package.Path)
	if _, err := os.Stat(d); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("ensuring '%s' is a valid directory: %w", d, err)
		}
	}

	w, err := os.Create(conf.Package.Path)
	if err != nil {
		return err
	}
	defer w.Close()

	if err := t.Execute(w, conf); err != nil {
		return err
	}

	return nil
}
