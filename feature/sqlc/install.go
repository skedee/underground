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
		srcFile := filepath.Join("feature/sqlc", "makefile")
		destFile := filepath.Join(flags.Project, "makefile")
		os.MergeEmbedFile(content, makefile.Keywords, srcFile, destFile)

		os.Mkdir(os.GetPath(flags.Project, "sqlc/db"))
		os.Mkdir(os.GetPath(flags.Project, "sqlc/query"))

		fileNames := []string{"sqlc.yaml", "README.md"}
		srcDir := "feature/sqlc"
		destDir := filepath.Join(flags.Project, "sqlc")
		os.CopyEmbedFiles(content, fileNames, srcDir, destDir)
	}
}
