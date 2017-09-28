package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func main() {
	accessid := os.Getenv("s3accessid")
	accesskey := os.Getenv("secretacceskey")
	token := ""
	cred := credentials.NewStaticCredentials(accessid, accesskey, token)
	config := aws.NewConfig()
	config.Credentials = cred
	config.Region = aws.String("us-east-2")
	svc := s3.New(session.New(config))
	list, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String("awsassign")})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}
