package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Function for reference of remove unnecessary files, mostly for .DS_Store.
// Please note that, the unnecessary file definition is only set in hardcoded.
func isUnnecessaryFiles(filename string) bool {
	unnecessaryFiles := []string{".DS_Store"}

	for _, item := range unnecessaryFiles {
		if item == filename {
			return true
		}
	}

	return false
}

// Get all absolute paths in rootDir.
// If any unnecessary files are found, this function will remove these files.
func getEmptyDir(rootDir string) []string {
	var directories []string

	// Remove all unnecessary_files and get directory structure
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {

		// Unnecessary files Detected
		if isUnnecessaryFiles(info.Name()) {
			os.Remove(path)
			fmt.Printf("Unnecessary File Detected and Deleted: %s\n", path)
		}

		// Record directory
		if info.IsDir() {
			directories = append(directories, path)
		}

		return nil
	})

	return directories
}

// Remove all empty directories in the root directory.
func RemoveAllEmptyDir(rootDir string) {
	fmt.Printf("Start checking Empty Directory. Getting All Directory...\n")

	// Get directory list
	directories := getEmptyDir(rootDir)

	// Loop through directories
	for len(directories) > 0 {
		// Wait list to check the directory again
		waitingList := make(map[string]struct{})

		for _, directory := range directories {
			// Get file count in current checking directory
			files, _ := os.ReadDir(directory)

			// Remove directory if it is empty
			if len(files) == 0 {
				fmt.Printf("Empty Directory detected: %s\n", directory)
				err := os.Remove(directory)
				if err != nil {
					fmt.Printf("Error when removing folder %s: %v\n", directory, err)
				} else {
					fmt.Printf("Successfully removed folder %s\n", directory)
				}

				// Put upper folder to wait list
				fmt.Printf("Upper Folder: %s Added to Wait list\n", filepath.Dir(directory))
				waitingList[filepath.Dir(directory)] = struct{}{}
			}
		}

		// Use Wait list to replace the dir_list to allow looping
		directories = make([]string, 0)
		for key := range waitingList {
			directories = append(directories, key)
		}

		fmt.Printf("Wait List: %s\n", directories)
		if len(directories) > 0 {
			fmt.Printf("Perform Search Again...\n\n")
		}
	}

	// Notify the process is completed
	fmt.Printf("Process Completed.\n")
	fmt.Printf("All Empty Directory removed.")
}

func main() {
	current_dir, _ := os.Getwd()
	RemoveAllEmptyDir(current_dir)
	fmt.Println("Press Enter to leave...")
	fmt.Scanln()
}
