package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"underground/flags"
)

//go:embed .makefile makefile env.sqitch README.md .gitignore
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

	fmt.Printf("Copied embedded file: %s\n", destinationFilePath)
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
		destinationFilePath := filepath.Join(dest, entry.Name())
		if err := os.WriteFile(destinationFilePath, fileContent, 0644); err != nil {
			return err
		}

		fmt.Printf("Copied embedded file: %s\n", destinationFilePath)
	}

	return nil
}

func mkdir(destDir string) {
	// Create the directory with 0755 permissions (read, write, and execute for owner, read, and execute for group and others)
	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory: ", err)
		return
	}

	fmt.Println("Directory created successfully:", destDir)
}

func main() {
	fileNames := []string{"makefile", "env.sqitch", "README.md", ".gitignore"}
	srcDir := ".makefile"

	flags.Parse()
	if flags.Project != "" {
		destDir := flags.Project
		// Write the content to the destination directory
		destinationFilePath := filepath.Join(destDir + "/" + srcDir)
		mkdir(destinationFilePath)

		for _, fileName := range fileNames {
			embedFileToDirectory(fileName, destinationFilePath)
		}

		if err := embedToDirectory(srcDir, destinationFilePath); err != nil {
			fmt.Printf("Error embedding files: %s", destinationFilePath)
		}
	}
}
