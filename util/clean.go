package util

import "os"

func Clean() {
	doClean(HEIC_DIR)
	doClean(JPEG_DIR)
}

func doClean(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		exitErrorf("failed to list dir %s", path)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			err := os.Remove(path + "/" + file.Name())
			if err != nil {
				exitErrorf("failed to remove %s", file.Name())
			}
		}
	}
}
