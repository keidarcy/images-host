package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	BUCKET_NAME = "super-food-gallery-2023"
	IMGIX_URL   = "https://zzw-food-gallery.imgix.net/"
)

type PageData struct {
	Title     string
	ImgOrigin string
	ImgNames  []string
}

func main() {
	file, err := os.Open("./public/index.tpl")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	htmlBytes, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	htmlString := string(htmlBytes)

	images := getImages()

	title := "zzw's food gallery 👩‍🍳 🍳"
	data := PageData{
		Title:     title,
		ImgOrigin: IMGIX_URL,
		ImgNames:  images,
	}
	tmpl := template.Must(template.New("html").Parse(htmlString))

	newFile, _ := os.Create("./public/index.html")

	err = tmpl.Execute(newFile, data)

	if err != nil {
		panic(err)
	}
}

func getImages() []string {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	if err != nil {
		exitErrorf("new s3 session failed")
	}

	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(BUCKET_NAME)})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", BUCKET_NAME, err)
	}

	images := []string{}
	for _, item := range resp.Contents {
		images = append(images, *item.Key)
	}
	sort.Strings(images)
	return images
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
