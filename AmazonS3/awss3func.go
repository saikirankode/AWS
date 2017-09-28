package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
)

type file struct {
	Filename, Bucketname, Objectname string
}

func uploadobject(res http.ResponseWriter, req *http.Request) {
	var newfile file
	params := req.URL.Query()
	newfile.Filename = params["filename"][0]
	newfile.Bucketname = params["bucketname"][0]
	newfile.Objectname = params["objectname"][0]
	fmt.Fprintln(res, newfile)
	accessid := "AKIAJGNZKBMDTCG3T7GA"
	accesskey := "j+aFsGQpf4+bCvCqFJ8hHKUF094688eMGRVgRVWQ"
	token := ""
	cred := credentials.NewStaticCredentials(accessid, accesskey, token)
	config := aws.NewConfig()
	config.Credentials = cred
	config.Region = aws.String("us-east-2")
	sess := session.Must(session.NewSession(config))
	upload := s3manager.NewUploader(sess)
	file := newfile.Filename
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	buckname := newfile.Bucketname
	objname := newfile.Objectname
	uploadinput := s3manager.UploadInput{Bucket: aws.String(buckname), Key: aws.String(objname), Body: f}
	up, err := upload.Upload(&uploadinput)
	if err != nil {
		fmt.Println("upload error", err)
	}
	fmt.Println(up, "done uploadong")
}
