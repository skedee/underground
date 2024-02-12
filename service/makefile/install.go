package makefile

import (
	"embed"
	"fmt"
	"path/filepath"
	"underground/flags"
	"underground/os"
)

func Install(content embed.FS) {
	destDir := filepath.Join(flags.Project, ".makefile")
	srcDir := "service/makefile/.makefile"
	os.Mkdir(destDir)

	if err := os.CopyEmbedToDirectory(content, srcDir, destDir); err != nil {
		fmt.Printf("Error embedding files: %s", destDir)
	}

	fileNames := []string{"makefile", "README.md"}
	srcDir = "service/makefile"
	os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)
}
