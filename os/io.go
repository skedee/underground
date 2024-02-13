package os

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InsertAfterKeyword(filename string, keyword string, lines []string) error {
	// Open the file for reading and writing
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var outputLines []string
	// Read lines until the keyword is found
	for scanner.Scan() {
		line := scanner.Text()
		outputLines = append(outputLines, line)
		if strings.Contains(line, keyword) {
			fmt.Printf("\nfound %s\n", keyword)
			// Insert the lines after the keyword
			outputLines = append(outputLines, lines...)
		}
	}
	// Write the updated content back to the file
	file.Seek(0, 0)
	file.Truncate(0)
	writer := bufio.NewWriter(file)
	for _, line := range outputLines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func InsertBeforeKeyword(filename string, keyword string, lines []string) error {
	// Open the file for reading
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a temporary file to write the modified content
	tmpfile, err := os.CreateTemp("", "temp.*.txt")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %w", err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Check if the line contains the keyword
		if strings.Contains(line, keyword) {
			// Write the lines to be inserted before the line with the keyword
			for _, l := range lines {
				_, err := tmpfile.WriteString(l + "\n")
				if err != nil {
					return fmt.Errorf("error writing to temporary file: %w", err)
				}
			}
		}
		_, err := tmpfile.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to temporary file: %w", err)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Rename the temporary file to replace the original file
	err = os.Rename(tmpfile.Name(), filename)
	if err != nil {
		return fmt.Errorf("error renaming temporary file: %w", err)
	}

	return nil
}

func GrepLinesBetweenKeywords(filename, startKeyword, endKeyword string) ([]string, error) {
	// Open the file for reading
	file, err := os.Open(filename)
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

func Replace(filename, oldKeyword, newKeyword string) error {
	// Read the content of the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Replace all instances of the old keyword with the new keyword
	newContent := strings.ReplaceAll(string(content), oldKeyword, newKeyword)

	// Write the updated content back to the file
	err = os.WriteFile(filename, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func MergeFiles(keywords []string, sourceFile, destFile string) {
	for _, keyword := range keywords {
		lines, _ := GrepLinesBetweenKeywords(sourceFile, keyword+"@", keyword+"$")
		fmt.Printf("AAAAAAA %s: %s: %s\n", sourceFile, keyword, lines)
		if len(lines) > 0 {
			InsertAfterKeyword(destFile, keyword+"@", lines)
		}
	}
}
