package util

import (
	"fmt"
	"os"
)

var (
	DOWNLOAD_DIR = "Downloads"
	HEIC_DIR     = "./heic-images"
	JPEG_DIR     = "./jpeg-images"
	IMGIX_URL    = os.Getenv("IMGIX_URL")
	BUCKET_NAME  = "super-food-gallery-2023"
	AWS_REGION   = "ap-northeast-1"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
