package commons

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// List files by pattern and creation time
func listFilesByPatternAndCreationTime(directoryPath, searchPattern string) ([]os.FileInfo, error) {

	// Read the directory
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	// Filter files by search pattern
	var matchingFiles []os.FileInfo
	for _, file := range files {
		if strings.HasPrefix(file.Name(), searchPattern) {
			matchingFiles = append(matchingFiles, file)
		}
	}

	// Sort files by creation time in descending order
	sort.Slice(matchingFiles, func(i, j int) bool {
		return matchingFiles[i].ModTime().After(matchingFiles[j].ModTime())
	})

	// Return the matching files - vai porra!
	return matchingFiles, nil
}
