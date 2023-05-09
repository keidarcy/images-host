package util

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Move() {

	home, err := os.UserHomeDir()
	if err != nil {
		exitErrorf("failed to get home dir")
	}
	sourceDir := fmt.Sprintf("%s/%s", home, DOWNLOAD_DIR)

	files, err := os.ReadDir(sourceDir)
	if err != nil {
		exitErrorf("failed to read download dir")
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			name := strings.ToLower(file.Name())
			if strings.Contains(name, "heic") {
				doMove(name, sourceDir)
			}
		}
	}

}

func doMove(name string, source string) {
	targetName := HEIC_DIR + "/" + name
	sourceName := source + "/" + name
	sourceFile, err := os.Open(sourceName)
	if err != nil {
		exitErrorf("failed to open source file %s", sourceName)
	}

	targetFile, err := os.Create(targetName)
	if err != nil {
		exitErrorf("failed to create targe file %s", targetName)
	}

	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		exitErrorf("failed to copy file %s to file %s", sourceName, targetName)
	}

	err = os.Remove(sourceName)
	if err != nil {
		exitErrorf("failed to remove file %s", sourceName)
	}

	fmt.Printf("file %s moved to %s\n", sourceName, targetName)
}
