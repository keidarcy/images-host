package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type Item struct {
	Id        string
	CreatedAt string
	UpdatedAt string
	Name      string
	URL       string
	Deleted   bool
}

func NewItem(url string, day time.Duration) *Item {
	id := uuid.New()
	// createdAt := time.Now().Format(time.RFC3339)
	createdAt := time.Now().Add(-24 * day * time.Hour).Format(time.RFC3339)

	return &Item{
		Id:        id.String(),
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		URL:       url,
		Deleted:   false,
		Name:      " ",
	}
}

func read(sess *session.Session) {
	svc := dynamodb.New(sess)
	tableName := "food-gallery"

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	var items []Item

	err := svc.ScanPages(input, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		for _, item := range page.Items {
			fmt.Printf("item: %v\n", item)
			var itemData Item
			err := dynamodbattribute.UnmarshalMap(item, &itemData)
			if err != nil {
				fmt.Println("Error unmarshaling item: ", err)
				return false
			}
			items = append(items, itemData)
		}
		return true
	})

	if err != nil {
		fmt.Println("Error scanning table: ", err)
		return
	}

	for _, item := range items {
		fmt.Println(item.Name)
	}
}

func create(sess *session.Session, urls []string, tableName string) error {
	for index, url := range urls {
		day := time.Duration(len(urls) + 1 - index)
		err := doPutItem(sess, url, tableName, day)
		if err != nil {
			return err
		}
	}
	return nil
}

func doPutItem(sess *session.Session, url, tableName string, day time.Duration) error {
	svc := dynamodb.New(sess)
	item := NewItem(url, day)

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err == nil {
		fmt.Printf("%s put item success\n", url)
	}
	return err
}
