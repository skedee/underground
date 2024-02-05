package sqitch

import (
	"embed"
	"fmt"
	"path/filepath"
	"underground/flags"
	"underground/os"
)

func Install(content embed.FS) {
	destDir := filepath.Join(flags.Project, ".makefile")
	srcDir := "service/sqitch/.makefile"

	os.Mkdir(destDir)

	if err := os.EmbedToDirectory(content, srcDir, destDir); err != nil {
		fmt.Printf("Error embedding files: %s", destDir)
	}

	fileNames := []string{"makefile", "env.sqitch", "README.md"}
	srcDir = "service/sqitch"

	for _, fileName := range fileNames {
		src := filepath.Join(srcDir, fileName)
		destDir = filepath.Join(flags.Project, fileName)
		os.EmbedFileToDirectory(content, src, destDir)
	}
}
