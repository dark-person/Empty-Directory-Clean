package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func isUnnecessaryFiles(filename string) bool {
	// Function for reference of remove unnecessary files, mostly for .DS_Store
	// The unnecessary file only set in hardcoded

	unnecessary_files := []string{".DS_Store"}

	for _, item := range unnecessary_files {
		if item == filename {
			return true
		}
	}

	return false
}

func RemoveAllEmptyDir2(root_dir string) {
	fmt.Printf("Start checking Empty Directory. Getting All Directory...\n")
	var dir_list []string

	// Remove all unnecessary_files and get directory structure
	filepath.Walk(root_dir, func(path string, info os.FileInfo, err error) error {
		//fmt.Printf("Current Walking: %s\n", path)
		if isUnnecessaryFiles(info.Name()) {
			// Unnecessary files Detected
			os.Remove(path)
			fmt.Printf("Unneccessary File Detected and Deleted: %s\n", path)
		}
		if info.IsDir() {
			dir_list = append(dir_list, path)
		}
		return nil
	})

	for len(dir_list) > 0 {
		// Wait list to check the directory again
		wait_list := make(map[string]struct{})
		for _, directory := range dir_list {
			//fmt.Printf("Current Checking Directory: %s\n", directory)

			// Get file count in current checking directory
			files, _ := os.ReadDir(directory)
			//fmt.Printf("Current File Count: %d\n", len(files))

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
				wait_list[filepath.Dir(directory)] = struct{}{}
			}
		}

		// Use Wait list to replace the dir_list to allow looping
		dir_list = make([]string, 0)
		for key := range wait_list {
			dir_list = append(dir_list, key)
		}

		fmt.Printf("Wait List: %s\n", dir_list)
		if len(dir_list) > 0 {
			fmt.Printf("Perform Search Again...\n\n")
		}

	}

	// Notify the process is completed
	fmt.Printf("Process Completed.\n")
	fmt.Printf("All Empty Directory removed.")
}

func main() {
	current_dir, _ := os.Getwd()
	RemoveAllEmptyDir2(current_dir)
	fmt.Println("Press Enter to leave...")
	fmt.Scanln()
}
