package db

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

func NewItem(url string) Item {
	id := uuid.New()
	createdAt := time.Now().Format(time.RFC3339)
	return Item{
		Id:        id.String(),
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		URL:       url,
		Deleted:   false,
	}
}

func GetItems(sess *session.Session, tableName string) ([]*Item, error) {
	svc := dynamodb.New(sess)

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	var items []*Item

	// result, _ := svc.Scan(input)
	// fmt.Printf("result: %v\n", result)

	count := 0
	err := svc.ScanPages(input, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		for _, item := range page.Items {
			count++
			if count > 2 {
				return false
			}
			var itemData Item
			err := dynamodbattribute.UnmarshalMap(item, &itemData)
			if err != nil {
				return false
			}
			if itemData.Deleted {
				return false
			}
			items = append(items, &itemData)
		}
		return true
	})

	if err != nil {
		fmt.Println("Error scanning table: ", err)
		return nil, err
	}
	return items, nil
	// for _, item := range items {
	// 	fmt.Println(item.Name)
	// }
}

func create(sess *session.Session, urls []string, tableName string) error {
	for _, url := range urls {
		err := doPutItem(sess, url, tableName)
		if err != nil {
			return err
		}
	}
	return nil
}

func doPutItem(sess *session.Session, url, tableName string) error {
	svc := dynamodb.New(sess)
	item := NewItem(url)

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
