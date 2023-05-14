package util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DOWNLOAD_DIR    = "Downloads"
	HEIC_DIR        = "./heic-images"
	JPEG_DIR        = "./jpeg-images"
	IMGIX_URL       = ""
	AWS_BUCKET_NAME = ""
	AWS_REGION      = ""
	TITLE           = "zzw's food gallery 👩‍🍳 🍳"
	AWS_TABLE_NAME  = ""
)

func init() {
	err := godotenv.Load()

	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file")
	}
	IMGIX_URL = os.Getenv("IMGIX_URL")
	AWS_BUCKET_NAME = os.Getenv("AWS_BUCKET_NAME")
	AWS_REGION = os.Getenv("AWS_REGION")
	AWS_TABLE_NAME = os.Getenv("AWS_TABLE_NAME")
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
