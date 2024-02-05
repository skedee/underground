package main

import (
	"embed"
	"underground/flags"
	"underground/service/sqitch"
)

//go:embed service/sqitch/.makefile service/sqitch/makefile service/sqitch/env.sqitch service/sqitch/README.md
var contentSqitch embed.FS

func main() {

	flags.Parse()
	if flags.Project != "" {
		sqitch.Install(contentSqitch)
	}
}
