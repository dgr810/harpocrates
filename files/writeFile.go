package files

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// WriteFile will write some string data to a file
func WriteFile(dirPath string, fileName string, content string) {
	fileName = fixFileName(fileName)
	path := filepath.Join(dirPath, fileName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0700)
		if err != nil {
			fmt.Printf("Unable to create dir at path '%s': %v\n", dirPath, err)
			os.Exit(1)
		}
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("An error happened while trying to open file %s: %s\n", path, err)
		os.Exit(1)
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		fmt.Printf("Unable to write to file '%s': %v\n", path, err)
		f.Close()
		os.Exit(1)
	}
}

func fixFileName(name string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9.]+")
	fileName := reg.ReplaceAllString(name, "_")

	return fileName
}
