package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"underground/flags"

	"github.com/skedee/leicester/log"
)

//go:embed .makefile makefile env.sqitch.example README.md .gitignore
var content embed.FS

func embedFileToDirectory(filename, destDir string) {
	// Read the content of the embedded file
	data, err := content.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading embedded file:", err)
		return
	}

	// Destination file path
	destinationFilePath := filepath.Join(destDir, filename)

	// Write the content to the destination file
	err = os.WriteFile(destinationFilePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to destination file:", err)
		return
	}

	fmt.Printf("Copied embedded file to: %s\n", destinationFilePath)
}

func embedToDirectory(src, dest string) error {
	// Ensure the destination directory exists
	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	// List the contents of the embedded directory
	dirContents, err := fs.ReadDir(content, src)
	if err != nil {
		return err
	}

	for _, entry := range dirContents {
		// Read the content of each file in the embedded directory
		fileContent, err := fs.ReadFile(content, filepath.Join(src, entry.Name()))
		if err != nil {
			return err
		}

		// Write the content to the destination directory
		destFilePath := filepath.Join(dest, entry.Name())
		if err := os.WriteFile(destFilePath, fileContent, 0644); err != nil {
			return err
		}

		fmt.Printf("File %s written to %s\n", entry.Name(), destFilePath)
	}

	return nil
}

func mkdir(destDir string) {
	// Create the directory with 0755 permissions (read, write, and execute for owner, read, and execute for group and others)
	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	fmt.Println("Directory created successfully at", destDir)
}

func main() {
	fileNames := []string{"makefile", "env.sqitch.example", "README.md", ".gitignore"}
	srcDir := ".makefile"

	flags.Parse()
	if flags.Project != "" {
		destDir := flags.Project
		mkdir(destDir + "/" + srcDir)

		for _, fileName := range fileNames {
			embedFileToDirectory(fileName, destDir)
		}

		if err := embedToDirectory(srcDir, destDir+"/"+srcDir); err != nil {
			log.Logger.Fatalf("Error embedding files: %s", destDir+"/"+srcDir)
		}
	}
}
