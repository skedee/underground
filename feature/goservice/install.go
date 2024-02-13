package goservice

import (
	"embed"
	"path/filepath"
	"underground/feature/makefile"
	"underground/flags"
	"underground/os"
)

func Install(content embed.FS) {
	if flags.GoService {
		fileNames := []string{"go.mod.txt", "go.sum.txt"}
		srcDir := "feature/goservice"
		os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)
		os.Rename(filepath.Join(flags.Project, "go.mod.txt"), filepath.Join(flags.Project, "go.mod"))
		os.Rename(filepath.Join(flags.Project, "go.sum.txt"), filepath.Join(flags.Project, "go.sum"))
		os.Replace(filepath.Join(flags.Project, "go.mod"), "@@project-name", flags.Project)

		srcFile := filepath.Join("feature/goservice", "makefile")
		destFile := filepath.Join(flags.Project, "makefile")
		os.MergeFiles(makefile.Keywords, srcFile, destFile)
	}
}
