package main

import (
	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/service/s3"

	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	//client:= s3.New(nil)
	accessid := os.Getenv("s3accessid")
	accesskey := os.Getenv("secretacceskey")
	token := ""
	cred := credentials.NewStaticCredentials(accessid, accesskey, token)
	config := aws.NewConfig()
	config.Credentials = cred
	config.Region = aws.String("us-east-2")
	sess := session.Must(session.NewSession(config))
	upload := s3manager.NewUploader(sess)
	f, err := os.Open("file")
	if err != nil {
		fmt.Println(err)
	}
	uploadinput := s3manager.UploadInput{Bucket: aws.String("awsassign"),Key:aws.String("a"), Body: f}
	up, err := upload.Upload(&uploadinput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(up,"done uploadong")

}
