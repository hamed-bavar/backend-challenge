package utils

import (
	"challenge/lib/errors"
	"challenge/lib/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

func GetDbClient() (*dynamo.DB, *errors.AppError) {
	credential := credentials.NewStaticCredentials(os.Getenv("ACCESS_TOKEN"), os.Getenv("SECRET_KEY"), "")
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2"), Credentials: credential})
	if err != nil {
		logger.Error("Error while creating db client: " + err.Error())
		return nil, errors.InternalServerError("Internal Server Error")
	}
	dbClient := dynamo.New(sess)
	return dbClient, nil
}
