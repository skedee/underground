package os

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func EmbedGrepLinesBetweenKeywords(content embed.FS, filename, startKeyword, endKeyword string) ([]string, error) {
	// Open the file for reading
	file, err := content.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	capturing := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip lines containing the start keyword
		if strings.Contains(line, startKeyword) {
			capturing = true
			continue
		}

		// If capturing is true, add the line to the result
		if capturing {
			// If the line contains the end keyword, stop capturing
			if strings.Contains(line, endKeyword) {
				break
			}
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func EmbedAppendToFile(content embed.FS, sourcePath, destinationPath string) error {
	// Read the content of the embed source
	sourceContent, err := content.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	// Open the destination file in append mode, or create the file if it doesn't exist
	destinationFile, err := os.OpenFile(destinationPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Write the content of the source file to the end of the destination file
	_, err = destinationFile.Write(sourceContent)
	if err != nil {
		return err
	}

	return nil
}

func CopyEmbedFiles(content embed.FS, fileNames []string, srcDir, destPath string) {
	for _, fileName := range fileNames {
		src := filepath.Join(srcDir, fileName) // need relative path for embed resource
		destDir := GetPath(destPath, fileName)
		CopyEmbedFileToDirectory(content, src, destDir)
	}
}

func CopyEmbedFileToDirectory(content embed.FS, src, dest string) {
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

func CopyEmbedToDirectory(content embed.FS, src, dest string) error {
	// Ensure the destination directory exists
	if err := Mkdir(dest); err != nil {
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

func GrepEmbedLinesBetweenKeywords(content embed.FS, filename, startKeyword, endKeyword string) ([]string, error) {
	// Open the file for reading
	file, err := content.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	capturing := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip lines containing the start keyword
		if strings.Contains(line, startKeyword) {
			capturing = true
			continue
		}

		// If capturing is true, add the line to the result
		if capturing {
			// If the line contains the end keyword, stop capturing
			if strings.Contains(line, endKeyword) {
				break
			}
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func MergeEmbedFile(content embed.FS, keywords []string, sourceFile, destFile string) {
	for _, keyword := range keywords {
		lines, _ := GrepEmbedLinesBetweenKeywords(content, sourceFile, keyword+"@", keyword+"$")
		if len(lines) > 0 {
			InsertAfterKeyword(destFile, keyword+"@", lines)
		}
	}
}
