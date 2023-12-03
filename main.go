package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var (
	AWS_REGION      = ""
	AWS_BUCKET_NAME = ""
	Table_Name      = "food-gallery"
)

func init() {
	err := godotenv.Load()

	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file")
	}
	AWS_BUCKET_NAME = os.Getenv("AWS_BUCKET_NAME")
	AWS_REGION = os.Getenv("AWS_REGION")
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION),
	})
	if err != nil {
		exitErrorf("new s3 session failed")
	}

	images := getImages(sess)
	err = create(sess, images, Table_Name)

	if err != nil {
		exitErrorf("put item failed")
	}

}

func getImages(sess *session.Session) []string {
	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(AWS_BUCKET_NAME)})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", AWS_BUCKET_NAME, err)
	}

	images := []string{}
	for _, item := range resp.Contents {
		images = append(images, *item.Key)
	}
	sort.Strings(images)
	fmt.Printf("images: %v\n", images[:3])
	return images
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
