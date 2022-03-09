package utils

import (
	"challenge/lib/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func GetDbClient() (*dynamodb.DynamoDB, *errors.AppError) {
	credential := credentials.NewStaticCredentials(os.Getenv("ACCESS_TOKEN"), os.Getenv("SECRET_KEY"), "")
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2"), Credentials: credential})
	if err != nil {
		return nil, errors.InternalServerError("Internal Server Error")
	}
	dbClient := dynamodb.New(sess)
	return dbClient, nil
}
