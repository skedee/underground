package sqitch

import (
	"embed"
	"fmt"
	"path/filepath"
	"underground/feature/makefile"
	"underground/flags"
	"underground/os"
)

func Install(content embed.FS) {
	if flags.Sqitch {
		destDir := filepath.Join(flags.Project, ".makefile")
		srcDir := "feature/sqitch/.makefile"

		if err := os.CopyEmbedToDirectory(content, srcDir, destDir); err != nil {
			fmt.Printf("Error embedding files: %s", destDir)
		}

		fileNames := []string{"env.sqitch", "README.md"}
		srcDir = "feature/sqitch"
		os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)

		srcFile := filepath.Join("feature/sqitch", "makefile")
		destFile := filepath.Join(flags.Project, "makefile")
		os.MergeEmbedFile(content, makefile.Keywords, srcFile, destFile)
	}
}
