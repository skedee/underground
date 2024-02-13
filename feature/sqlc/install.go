package sqlc

import (
	"embed"
	"path/filepath"
	"underground/feature/makefile"
	"underground/flags"
	"underground/os"
)

func Install(content embed.FS) {
	if flags.Sqlc {
		fileNames := []string{"sqlc.yaml", "README.md"}
		srcDir := "feature/sqlc"
		os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)

		srcFile := filepath.Join("feature/sqlc", "makefile")
		destFile := filepath.Join(flags.Project, "makefile")
		os.MergeEmbedFile(content, makefile.Keywords, srcFile, destFile)
	}
}
