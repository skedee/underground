package makefile

import (
	"embed"
	"path/filepath"
	"underground/flags"
	"underground/os"
)

var (
	Keywords = []string{"@include-pre", "@include", "@include-post", "@target"}
)

func Install(content embed.FS) {
	destDir := filepath.Join(flags.Project, ".makefile")
	srcDir := "feature/makefile/.makefile"
	os.Mkdir(destDir)
	os.CopyEmbedToDirectory(content, srcDir, destDir)

	fileNames := []string{"makefile", "README.md"}
	srcDir = "feature/makefile"
	os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)
}
