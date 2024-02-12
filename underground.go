package main

import (
	"embed"
	"underground/flags"
	"underground/os"
	"underground/service/makefile"
	"underground/service/sqitch"
)

//go:embed service/makefile/.makefile service/makefile/makefile service/makefile/README.md
var contentMakefile embed.FS

//go:embed service/sqitch/.makefile service/sqitch/makefile service/sqitch/env.sqitch service/sqitch/README.md
var contentSqitch embed.FS

//go:embed service/sqlc/makefile service/sqlc/sqlc.yaml service/sqlc/README.md
var contentSqlc embed.FS

func main() {
	flags.Parse()
	if flags.Project != "" {
		os.Mkdir(flags.Project)
		// create makefile
		makefile.Install(contentMakefile)
		// add sqitch files
		sqitch.Install(contentSqitch)
		// add sqlc files
		//sqlc.Install(contentSqlc)
	}
}
