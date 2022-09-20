package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Error Happened %s \n", err)
		os.Exit(1)
	}
}

func createDefaultFolders(target string) {
	defaultFolders := []string{"Music", "Videos", "Docs", "Images", "Others"}
	for _, folder := range defaultFolders {
		_, err := os.Stat(folder)
		if os.IsNotExist(err) {
			os.Mkdir(filepath.Join(target, folder), 0755)
		}
	}
}
func organizeFolder(targetFolder string) {
	// read the dir
	filesAndFolders, err := os.ReadDir(targetFolder)
	check(err)

	// to track how many files moved
	noOfFiles := 0

	for _, filesAndFolder := range filesAndFolders {
		// check for files
		if !filesAndFolder.IsDir() {
			fileInfo, err := filesAndFolder.Info()
			check(err)

			//get the file full path
			oldPath := filepath.Join(targetFolder, fileInfo.Name())
			fileExt := filepath.Ext(oldPath)

			// switch case to move files based on ext
			switch fileExt {
			case ".png", ".jpg", ".jpeg":
				newPath := filepath.Join(targetFolder, "Images", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".mp4", ".mov", ".avi", ".amv":
				newPath := filepath.Join(targetFolder, "Videos", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".pdf", ".docx", ".csv", ".xlsx":
				newPath := filepath.Join(targetFolder, "Docs", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".mp3", ".wav", ".aac":
				newPath := filepath.Join(targetFolder, "Music", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			default:
				newPath := filepath.Join(targetFolder, "Others", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			}
		}
	}

	// print how many files moved
	if noOfFiles > 0 {
		fmt.Printf("%v number of files moved\n", noOfFiles)
	} else {
		fmt.Printf("No files moved")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Which folder do you want to orginize? - ")
	scanner.Scan()

	input := scanner.Text()

	_, check := os.Stat(input)

	if os.IsNotExist(check) {
		fmt.Println("fodler doesn't exist")
		os.Exit(1)
	} else {
		createDefaultFolders(input)

		organizeFolder(input)
	}
}
