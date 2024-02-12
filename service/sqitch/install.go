package sqitch

import (
	"embed"
	"fmt"
	"path/filepath"
	"underground/flags"
	"underground/os"
)

func Install(content embed.FS) {
	if flags.Sqitch {
		// destDir := filepath.Join(flags.Project, ".makefile")
		// srcDir := "service/sqitch/.makefile"

		// if err := os.CopyEmbedToDirectory(content, srcDir, destDir); err != nil {
		// 	fmt.Printf("Error embedding files: %s", destDir)
		// }

		// fileNames := []string{"env.sqitch", "README.md"}
		// srcDir = "service/sqitch"
		// os.CopyEmbedFiles(content, fileNames, srcDir, flags.Project)

		keywords := []string{"@include-pre", "@include", "@include-post", "@target"}
		srcFile := filepath.Join("service/sqitch", "makefile")
		makefile := filepath.Join(flags.Project, "makefile")
		for _, keyword := range keywords {
			lines, _ := os.GrepLinesBetweenKeywords(srcFile, keyword+"@", keyword+"$")
			if len(lines) > 0 {
				fmt.Printf("aaaaaaa%s %s\n\n\n", keyword, lines)
				os.InsertStringsAfterKeyword(makefile, keyword, lines)
			}
		}
	}
}
