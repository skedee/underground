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

func AppendToFile(sourcePath, destinationPath string) error {
	// Read the content of the source file
	sourceContent, err := os.ReadFile(sourcePath)
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

func AppendTextBeforeFirstXYZ(filePath, textToAppend string) error {
	// Open the file for reading and writing
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line starts with "XYZ"
		if strings.HasPrefix(line, "XYZ") {
			// Move the file pointer back to the beginning of the line
			_, err := file.Seek(int64(-len(line)), os.SEEK_CUR)
			if err != nil {
				return err
			}

			// Write the new text before the matching line
			_, err = file.Write([]byte(textToAppend + line))
			if err != nil {
				return err
			}

			// Move the file pointer to the end of the line
			_, err = file.Seek(int64(len(textToAppend)), os.SEEK_CUR)
			if err != nil {
				return err
			}

			// Break out of the loop after the first match
			break
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("Text appended successfully.")
	return nil
}

func InsertFileAfterFirstXYZ(filePath, insertFilePath string) error {
	// Open the main file for reading and writing
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Open the insert file for reading
	insertFile, err := os.Open(insertFilePath)
	if err != nil {
		return err
	}
	defer insertFile.Close()

	// Create a temporary file to store the modified content
	tmpFilePath := filePath + ".tmp"
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	// Create a scanner to read the main file line by line
	scanner := bufio.NewScanner(file)

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()

		// Write the original line to the temporary file
		_, err := tmpFile.WriteString(line + "\n")
		if err != nil {
			return err
		}

		// Check if the line starts with "XYZ"
		if strings.HasPrefix(line, "XYZ") {
			// Copy the content of the insert file to the temporary file
			scannerInsert := bufio.NewScanner(insertFile)
			for scannerInsert.Scan() {
				insertLine := scannerInsert.Text()
				_, err := tmpFile.WriteString(insertLine + "\n")
				if err != nil {
					return err
				}
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	// Close and remove the original file
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Remove(filePath); err != nil {
		return err
	}

	// Rename the temporary file to the original file path
	if err := os.Rename(tmpFilePath, filePath); err != nil {
		return err
	}

	fmt.Println("File updated successfully.")
	return nil
}
