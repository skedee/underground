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
		fileNames := []string{"go.mod.txt", "go.sum.txt", "main.go"}
		srcDir := "feature/goservice"
		os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)
		os.Rename(filepath.Join(flags.Project, "go.mod.txt"), filepath.Join(flags.Project, "go.mod"))
		os.Rename(filepath.Join(flags.Project, "go.sum.txt"), filepath.Join(flags.Project, "go.sum"))
		os.Replace(filepath.Join(flags.Project, "go.mod"), "@@project-name", flags.Project)
		os.Replace(filepath.Join(flags.Project, "main.go"), "goservice", "main")

		srcFile := filepath.Join("feature/goservice", "makefile")
		destFile := os.GetPath(flags.Project, "makefile") // filepath.Join(flags.Project, "makefile")
		os.MergeEmbedFile(content, makefile.Keywords, srcFile, destFile)
	}
}