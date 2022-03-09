package utils

import (
	"challenge/lib/errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func GetDbClient() (*dynamodb.DynamoDB, *errors.AppError) {
	credential := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	fmt.Println(credential, "credentials")
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2"), Credentials: credential})
	if err != nil {
		fmt.Println("Error while creating session:", sess)
		return nil, errors.InternalServerError("Internal Server Error from dynamo db")
	}
	fmt.Println("Session created successfully", err)
	dbClient := dynamodb.New(sess)
	fmt.Println(dbClient, "db client")
	return dbClient, nil
}
