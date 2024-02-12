package os

import (
	"fmt"
	"os"
	"underground/flags"
)

func FileCreate(filename string) {
	Mkdir(flags.Project)
	// Check if the file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// File does not exist, create an empty file
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
	}
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