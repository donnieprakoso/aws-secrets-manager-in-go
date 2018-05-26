package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetSecret(name string) (*secretsmanager.GetSecretValueOutput, error) {

	// This example assumes that you're connecting to ap-southeast-1 region
	// For a full list of endpoints, you can refer to this site -> https://godoc.org/github.com/aws/aws-sdk-go/aws/endpoints
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApSoutheast1RegionID),
	}))

	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		/*
			 To address specific error, you can import this package:
				"github.com/aws/aws-sdk-go/aws/awserr"

			and use this example:

			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case secretsmanager.ErrCodeResourceNotFoundException:
					fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
				case secretsmanager.ErrCodeInvalidParameterException:
					fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
				case secretsmanager.ErrCodeInvalidRequestException:
					fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
				case secretsmanager.ErrCodeDecryptionFailure:
					fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
				case secretsmanager.ErrCodeInternalServiceError:
					fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return nil, err
		*/
	}

	return result, nil
}

func main() {
	secretName := "YOUR AWS SECRETS MANAGER NAME HERE"
	secret, _ := GetSecret(secretName)
	fmt.Println(secret)
}
