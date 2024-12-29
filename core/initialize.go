package core

import (
	"fmt"
	"io/fs"
	"os"
)

// Initialize creates directories programmatically from a list of paths
func Initialize() {
	// Retrieve the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error occurred while retrieving user home directory:", err)
		return
	}

	// Define directories to create
	directories := []string{
		fmt.Sprintf("%s/.planecdn", homeDir),
		fmt.Sprintf("%s/.planecdn/files", homeDir),
		fmt.Sprintf("%s/.planecdn/logs", homeDir), // Add more directories as needed
	}

	// Create each directory
	for _, dir := range directories {
		exists, err := dirExists(dir)
		if err != nil {
			fmt.Printf("Error while checking if directory %s exists: %v\n", dir, err)
			continue
		}

		if !exists {
			err := os.Mkdir(dir, fs.ModePerm)
			if err != nil {
				fmt.Printf("Error while creating directory %s: %v\n", dir, err)
			} else {
				fmt.Printf("Created directory: %s\n", dir)
			}
		} else {
			fmt.Printf("Directory already exists: %s\n", dir)
		}
	}
}

// dirExists checks if a directory exists
func dirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}
