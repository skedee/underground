package makefile

import (
	"embed"
	"fmt"
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

	if err := os.CopyEmbedToDirectory(content, srcDir, destDir); err != nil {
		fmt.Printf("Error embedding files: %s", destDir)
	}

	fileNames := []string{"makefile", "README.md"}
	srcDir = "feature/makefile"
	os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)
}
