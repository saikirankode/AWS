package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	accessid := "AKIAJGNZKBMDTCG3T7GA"
	accesskey := "j+aFsGQpf4+bCvCqFJ8hHKUF094688eMGRVgRVWQ"
	token := ""
	cred := credentials.NewStaticCredentials(accessid, accesskey, token)
	config := aws.NewConfig()
	config.Credentials = cred
	config.Region = aws.String("us-east-2")
	svc := s3.New((session.New(config)))
	deleteout, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String("awsassign"), Key: aws.String("a")})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(deleteout)
}
