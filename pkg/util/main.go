package util

import (
	"fmt"
	"os"
)

const (
	DOWNLOAD_DIR = "Downloads"
	HEIC_DIR     = "./heic-images"
	JPEG_DIR     = "./jpeg-images"
	BUCKET_NAME  = "super-food-gallery-2023"
	IMGIX_URL    = "https://zzw-food-gallery.imgix.net/"
	AWS_REGION   = "ap-northeast-1"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
