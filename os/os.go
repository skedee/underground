package os

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func EmbedFileToDirectory(content embed.FS, src, dest string) {
	// Read the content of the embedded file
	data, err := content.ReadFile(src)
	if err != nil {
		fmt.Println("Error reading embedded file:", err)
		return
	}

	// Write the content to the destination file
	err = os.WriteFile(dest, data, 0644)
	if err != nil {
		fmt.Println("Error writing to destination file:", err)
		return
	}

	fmt.Printf("Copied embedded file: %s\n", dest)
}

func EmbedToDirectory(content embed.FS, src, dest string) error {
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

func Mkdir(destDir string) {
	// Create the directory with 0755 permissions (read, write, and execute for owner, read, and execute for group and others)
	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory: ", err)
		return
	}

	fmt.Println("Directory created successfully:", destDir)
}
