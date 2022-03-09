package utils

import (
	"challenge/lib/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDbClient() (*dynamodb.DynamoDB, *errors.AppError) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		return nil, errors.InternalServerError("Internal Server Error from dynamo db")
	}
	dbClient := dynamodb.New(sess)
	return dbClient, nil
}
