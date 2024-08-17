package main

import (
	"empty-directory-clean/config"
	"fmt"
	"os"
	"path/filepath"
)

// Check if filename is matched in removal config filename list.
//
// Please note that developer SHOULD never put absolute paths in filename params,
// that params is for filename only, e.g. "foo.txt", but not "abc/foo.txt".
func isFileToRemove(c *config.RemovalConfig, filename string) bool {
	for _, item := range c.Files {
		if item == filename {
			return true
		}
	}

	return false
}

// Check if file has extension that matched in removal config filename list.
//
// Please note that developer SHOULD never put absolute paths in filename params,
// that params is for filename only, e.g. "foo.txt", but not "abc/foo.txt".
func isExtToRemove(c *config.RemovalConfig, filename string) bool {
	for _, item := range c.Extension {
		if filepath.Ext(filename) == item {
			return true
		}
	}

	return false
}

// Get all absolute paths in rootDir.
// If any unnecessary files are found, this function will remove these files.
func getAllDirectories(c *config.RemovalConfig, rootDir string) []string {
	var directories []string

	// Remove all unnecessary_files and get directory structure
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {

		// Unnecessary files Detected
		if isFileToRemove(c, info.Name()) {
			os.Remove(path)
			fmt.Printf("Detected and Deleted matched filename: %s\n", path)
		}

		// Detect file extensions
		if isExtToRemove(c, info.Name()) {
			os.Remove(path)
			fmt.Printf("Detected and Deleted matched extension: %s\n", path)
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
func RemoveAllEmptyDir(c *config.RemovalConfig, rootDir string) {
	fmt.Printf("Start checking Empty Directory. Getting All Directory...\n")

	// Get directory list
	directories := getAllDirectories(c, rootDir)

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
	fmt.Printf("All Empty Directory removed.\n")
}

func main() {
	current_dir, _ := os.Getwd()

	// Load configuration
	c, err := config.LoadYaml("config.yaml")
	if err != nil {
		fmt.Printf("No config file loaded, use default configuration.")
	} else {
		fmt.Printf("Config file loaded: " + c.String())
	}

	// Start remove process
	RemoveAllEmptyDir(c, current_dir)

	// Hold Terminal
	fmt.Println("Press Enter to leave...")
	fmt.Scanln()
}
