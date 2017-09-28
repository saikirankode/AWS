package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	sess := session.Must(session.NewSession(config))
	downloader := s3manager.NewDownloader(sess)
	f, err := os.Create("uploadfile")
	if err != nil {
		fmt.Println("cannot create a uploadfile", err)
	}
	dow, err := downloader.Download(f, &s3.GetObjectInput{Bucket: aws.String("awsassign"), Key: aws.String("a"),})
	if err != nil {
		fmt.Println("Error downloading File", err)
	}
	fmt.Println("new uploadfile download", dow)
}
