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
	// copy .makefile/
	destDir := filepath.Join(flags.Project, ".makefile")
	srcDir := "feature/makefile/.makefile"
	os.CopyEmbedToDirectory(content, srcDir, destDir)

	// copy docsify/
	destDir = flags.Project
	srcDir = "feature/makefile/docsify"
	os.CopyEmbedToDirectory(content, srcDir, destDir)
	os.Replace(os.GetPath(flags.Project, "index.html"), "@@project-name@@", flags.Project)

	fileNames := []string{"makefile", "README.md"}
	srcDir = "feature/makefile"
	os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)
}
