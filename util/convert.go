package util

import (
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jdeng/goheif"
)

func Convert() {
	heics, err := os.ReadDir(HEIC_DIR)

	if err != nil {
		log.Fatal("cannot load heic images")
	}

	for _, heic := range heics {
		if heic.IsDir() {
			fmt.Printf("%s is a directory", heic.Name())
		} else {
			name := strings.ToLower(heic.Name())
			if !strings.HasSuffix(name, "heic") {
				continue
			}
			doConvert(name[:len(name)-5])
		}
	}
}

func doConvert(name string) {
	sourcePath := fmt.Sprintf("%s/%s.HEIC", HEIC_DIR, name)
	targetPath := fmt.Sprintf("%s/%s.jpeg", JPEG_DIR, name)

	fmt.Printf("name: %v\n", name)
	fmt.Printf("sourcePath: %v\n", sourcePath)
	fmt.Printf("targetPath: %v\n", targetPath)

	fi, err := os.Open(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	img, err := goheif.Decode(fi)
	if err != nil {
		log.Fatalf("Failed to parse %s: %v\n", sourcePath, err)
	}

	fo, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0644)
	fmt.Printf("fo: %v\n", fo)
	if err != nil {
		log.Fatalf("Failed to create output file %s: %v\n", targetPath, err)
	}
	defer fo.Close()

	w, _ := newWriterExif(fo, nil)
	err = jpeg.Encode(w, img, nil)

	if err != nil {
		log.Fatalf("Failed to encode %s: %v\n", targetPath, err)
	}

	log.Printf("Convert %s to %s successfully\n", sourcePath, targetPath)
}

func newWriterExif(w io.Writer, exif []byte) (io.Writer, error) {
	writer := &writerSkipper{w, 2}
	soi := []byte{0xff, 0xd8}
	if _, err := w.Write(soi); err != nil {
		return nil, err
	}
	if exif != nil {
		app1Marker := 0xe1
		markerlen := 2 + len(exif)
		marker := []byte{0xff, uint8(app1Marker), uint8(markerlen >> 8), uint8(markerlen & 0xff)}
		if _, err := w.Write(marker); err != nil {
			return nil, err
		}

		if _, err := w.Write(exif); err != nil {
			return nil, err
		}
	}

	return writer, nil
}

// Skip Writer for exif writing
type writerSkipper struct {
	w           io.Writer
	bytesToSkip int
}

func (w *writerSkipper) Write(data []byte) (int, error) {
	if w.bytesToSkip <= 0 {
		return w.w.Write(data)
	}

	if dataLen := len(data); dataLen < w.bytesToSkip {
		w.bytesToSkip -= dataLen
		return dataLen, nil
	}

	if n, err := w.w.Write(data[w.bytesToSkip:]); err == nil {
		n += w.bytesToSkip
		w.bytesToSkip = 0
		return n, nil
	} else {
		return n, err
	}
}
