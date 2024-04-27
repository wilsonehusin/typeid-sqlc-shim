package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func errorLog(str string, args ...any) {
	color.New(color.FgHiRed, color.Bold).Fprint(os.Stderr, "[ERROR] ")
	fmt.Fprintf(os.Stderr, str, args...)
}

func warnLog(str string, args ...any) {
	color.New(color.FgYellow, color.Bold).Fprint(os.Stderr, "[WARN] ")
	fmt.Fprintf(os.Stderr, str, args...)
}

func infoLog(str string, args ...any) {
	color.New(color.FgHiGreen, color.Bold).Fprint(os.Stderr, "[INFO] ")
	fmt.Fprintf(os.Stderr, str, args...)
}
