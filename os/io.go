package os

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InsertStringsAfterKeyword(filename string, keyword string, lines []string) error {
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
		_, err := tmpfile.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to temporary file: %w", err)
		}
		// Check if the line contains the keyword
		if strings.Contains(line, keyword) {
			// Write the lines to be inserted after the line with the keyword
			for _, l := range lines {
				_, err := tmpfile.WriteString(l + "\n")
				if err != nil {
					return fmt.Errorf("error writing to temporary file: %w", err)
				}
			}
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

func InsertStringsBeforeKeyword(filename string, keyword string, lines []string) error {
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

func Sed(filename, arg string) error {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a temporary file to write the modified content
	tmpfile, err := os.CreateTemp("", "sqitch.conf.*.tmp")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %w", err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Perform the substitution
		line = strings.Replace(line, "@@registry", fmt.Sprintf("registry = sqitch_%s", arg), -1)
		// Write the modified line to the temporary file
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

func InsertFileBeforeKeyword(keyword, filename, insertFilename string) error {
	// Open the original file for reading
	file, err := os.Open(filename)
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

	// Open the file to be inserted
	insertFile, err := os.Open(insertFilename)
	if err != nil {
		return fmt.Errorf("error opening insert file: %w", err)
	}
	defer insertFile.Close()

	// Create a scanner to read the original file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Looking for: %s\n", keyword)
		// Check if the line contains the keyword
		if strings.Contains(line, keyword) {
			fmt.Printf("FOUND %s\n", keyword)
			// Write the content of the insert file before the line with the keyword
			insertScanner := bufio.NewScanner(insertFile)
			for insertScanner.Scan() {
				_, err := tmpfile.WriteString(insertScanner.Text() + "\n")
				if err != nil {
					return fmt.Errorf("error writing to temporary file: %w", err)
				}
			}
		}
		// Write the original line to the temporary file
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

func InsertFileAfterKeyword(keyword, filename, insertFilename string) error {
	// Open the original file for reading
	file, err := os.Open(filename)
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

	// Open the file to be inserted
	insertFile, err := os.Open(insertFilename)
	if err != nil {
		return fmt.Errorf("error opening insert file: %w", err)
	}
	defer insertFile.Close()

	// Create a scanner to read the original file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Write the original line to the temporary file
		_, err := tmpfile.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to temporary file: %w", err)
		}
		// Check if the line contains the keyword
		if strings.Contains(line, keyword) {
			// Write the content of the insert file after the line with the keyword
			insertScanner := bufio.NewScanner(insertFile)
			for insertScanner.Scan() {
				_, err := tmpfile.WriteString(insertScanner.Text() + "\n")
				if err != nil {
					return fmt.Errorf("error writing to temporary file: %w", err)
				}
			}
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
