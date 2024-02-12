package sqlc

import (
	"embed"
	"path/filepath"
	"underground/flags"
	"underground/os"
)

func Install(content embed.FS) {
	if flags.Sqlc {
		fileNames := []string{"sqlc.yaml", "README.md"}
		srcDir := "service/sqlc"
		os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)

		destMakefile := filepath.Join(flags.Project, "makefile")
		srcMakefile := filepath.Join(srcDir, "makefile")
		os.InsertFileBeforeKeyword("help:", destMakefile, srcMakefile)
	}
}
