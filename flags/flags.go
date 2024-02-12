package flags

import (
	"flag"
	"fmt"
	"os"
)

var (
	Project = ""
	Sqitch  = false
	Sqlc    = false
)

func Parse() {
	showHelp := false

	flag.BoolVar(&showHelp, "h", false, "Display usage information")
	flag.StringVar(&Project, "p", Project, "New underground project")
	flag.BoolVar(&Sqitch, "s", Sqitch, "Add Sqitch to project")
	flag.BoolVar(&Sqlc, "c", Sqlc, "Add Sqlc to project")

	flag.Parse() // Parse the flags

	// If the help flag is provided, show usage and exit
	// The default values are not printed for all flags when using flag.PrintDefaults() because the flag package
	// considers the zero values of the types when determining if a flag's default value should be displayed.
	// Only when the user provides a value different from the zero value, it will be displayed as the default.
	if showHelp {
		fmt.Println("Usage: Underground [options]")
		flag.PrintDefaults()
		os.Exit(0)
	}
}
