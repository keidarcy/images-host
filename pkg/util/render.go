package util

import (
	"io"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/keidarcy/zzw-food-gallery/pkg/db"
)

type Img struct {
	URL  string
	Name string
}

type PageData struct {
	Title     string
	ImgOrigin string
	Images    []Img
	BuildTime string
}

func Render() {

	file, err := os.Open("./public/index.tpl")
	if err != nil {
		exitErrorf("failed to open template", err)
	}
	defer file.Close()

	htmlBytes, err := io.ReadAll(file)

	if err != nil {
		exitErrorf("failed to read from template to bytes", err)
	}

	htmlString := string(htmlBytes)

	images, err := getImages()
	if err != nil {
		exitErrorf("failed to get images %v", err)
	}

	title := TITLE
	data := PageData{
		Title:     title,
		ImgOrigin: IMGIX_URL,
		Images:    images,
		BuildTime: time.Now().Format(time.RFC3339),
	}
	tmpl := template.Must(template.New("html").Parse(htmlString))

	newFile, _ := os.Create("./public/index.html")

	err = tmpl.Execute(newFile, data)

	if err != nil {
		exitErrorf("failed to execute template", err)
	}
}

// func getImages() []string {
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String(AWS_REGION),
// 	})

// 	if err != nil {
// 		exitErrorf("new s3 session failed")
// 	}

// 	svc := s3.New(sess)

// 	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(AWS_BUCKET_NAME)})
// 	if err != nil {
// 		exitErrorf("Unable to list items in bucket %q, %v", AWS_BUCKET_NAME, err)
// 	}

// 	images := []string{}
// 	for _, item := range resp.Contents {
// 		images = append(images, *item.Key)
// 	}
// 	sort.Strings(images)
// 	return images
// }

func getImages() ([]Img, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION),
	})
	if err != nil {
		log.Fatal("failed new session")
	}
	items, err := db.GetItems(sess, AWS_TABLE_NAME)
	imgs := []Img{}
	for _, item := range items {
		img := Img{
			Name: item.Name,
			URL:  item.URL,
		}
		imgs = append(imgs, img)
	}
	if err != nil {
		log.Fatal("failed get items")
	}

	return imgs, nil
}
