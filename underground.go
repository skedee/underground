package main

import (
	"embed"
	"underground/feature/goservice"
	"underground/feature/makefile"
	"underground/feature/sqitch"
	"underground/feature/sqlc"
	"underground/flags"
	"underground/os"
)

var (
	//go:embed feature/goservice/makefile feature/goservice/go.mod.txt feature/goservice/go.sum.txt
	contentGoService embed.FS

	//go:embed feature/makefile/.makefile feature/makefile/makefile feature/makefile/README.md
	contentMakefile embed.FS

	//go:embed feature/sqitch/.makefile feature/sqitch/makefile feature/sqitch/env.sqitch feature/sqitch/README.md
	contentSqitch embed.FS

	//go:embed feature/sqlc/makefile feature/sqlc/sqlc.yaml feature/sqlc/README.md
	contentSqlc embed.FS
)

func main() {
	flags.Parse()
	if flags.Project != "" {
		os.Mkdir(flags.Project)
		// // create makefile
		makefile.Install(contentMakefile)
		// add golang service files
		goservice.Install(contentGoService)
		// add sqitch files
		sqitch.Install(contentSqitch)
		// add sqlc files
		sqlc.Install(contentSqlc)
	}
}
