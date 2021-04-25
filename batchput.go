package batchidioms

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gocarina/gocsv"
	"log"
)

var dy = dynamodb.New(session.Must(session.NewSession(&aws.Config{
	Endpoint: aws.String("http://localhost:4566"),        // LocalStack
	Region:   aws.String(endpoints.ApNortheast1RegionID), // Tokyo Region
})))

func BatchWrite(ctx context.Context, writes []Forum) error {
	if len(writes) > 25 {
		return errors.New("batch write size is within 25 items")
	}

	items := make([]*dynamodb.WriteRequest, 0, len(writes))
	for _, v := range writes {
		av, _ := dynamodbattribute.MarshalMap(v) // エラーハンドリングは省略
		items = append(items, &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: av,
			},
		})
	}

	for len(items) > 0 {
		out, err := dy.BatchWriteItemWithContext(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				"forum": items,
			},
		})
		if err != nil {
			return fmt.Errorf("batch write to %s: %w", "music", err)
		}

		items = append(items[:0], out.UnprocessedItems["forum"]...) // 未処理のitemsがあれば再設定
	}

	return nil
}

func LoadForums() []Forum {
	data := `Name,Category
Amazon DynamoDB,Amazon Web Services
Amazon RDS,Amazon Web Services
Amazon Redshift,Amazon Web Services
Amazon ElastiCache,Amazon Web Services
`

	var forums []Forum
	if err := gocsv.UnmarshalBytes([]byte(data), &forums); err != nil {
		log.Fatal(err)
	}
	return forums
}
