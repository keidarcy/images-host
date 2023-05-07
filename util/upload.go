package util

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION)},
	)
	if err != nil {
		exitErrorf("failed to get new s3 session")
	}

	jpgs, err := os.ReadDir(JPEG_DIR)
	if err != nil {
		exitErrorf("failed to read dir %s, err: %v", JPEG_DIR, err)
	}

	uploader := s3manager.NewUploader(sess)
	for _, jpg := range jpgs {
		doUpload(uploader, jpg.Name())
	}
}

func doUpload(uploader *s3manager.Uploader, filename string) {

	file, err := os.Open(JPEG_DIR + "/" + filename)
	if err != nil {
		exitErrorf("failed to open file %s", filename)
	}
	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, BUCKET_NAME, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, BUCKET_NAME)
}
