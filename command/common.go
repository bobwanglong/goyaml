package command

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var CommandlineGlobalFlags GlobalFlags

// GlobalFlags is the global flags for the whole client.
type GlobalFlags struct {
	Server       string
	OutputFormat string
}

func ExitWithErrorf(format string, a ...interface{}) {
	ExitWithError(fmt.Errorf(format, a...))
}
func ExitWithError(err error) {
	if err != nil {
		color.New(color.FgRed).Fprint(os.Stderr, "Error: ")
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
