package utils

import (
	"challenge/lib/errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDbClient() (*dynamodb.DynamoDB, *errors.AppError) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		fmt.Println("Error while creating session:", err)
		return nil, errors.InternalServerError("Internal Server Error from dynamo db")
	}
	fmt.Println("Session created successfully", err)
	dbClient := dynamodb.New(sess)
	fmt.Println(dbClient, "db client")
	return dbClient, nil
}
