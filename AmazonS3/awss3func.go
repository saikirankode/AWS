package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
	"github.com/aws/aws-sdk-go/service/s3"
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
	accessid := os.Getenv("s3accessid")
	accesskey := os.Getenv("secretacceskey")
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

func Listobject(res http.ResponseWriter, req *http.Request) {
	var newfile file
	params := req.URL.Query()
	newfile.Bucketname=params["bucketname"][0]
	accessid := os.Getenv("s3accessid")
	accesskey := os.Getenv("secretacceskey")
	token := ""
	cred := credentials.NewStaticCredentials(accessid, accesskey, token)
	config := aws.NewConfig()
	config.Credentials = cred
	config.Region = aws.String("us-east-2")
	svc := s3.New(session.New(config))
	buckname := newfile.Bucketname
	list, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(buckname)})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}

func Downloadobject(res http.ResponseWriter, req *http.Request) {
	var newfile file
	params := req.URL.Query()
	newfile.Bucketname=params["bucketname"][0]
	newfile.Objectname=params["objectname"][0]
	fmt.Fprintln(res,newfile)
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
		fmt.Println("cannot create a file", err)
	}
	buckname := newfile.Bucketname
	objname := newfile.Objectname
	dow, err := downloader.Download(f, &s3.GetObjectInput{Bucket: aws.String(buckname), Key: aws.String(objname),})
	if err != nil {
		fmt.Println("Error downloading File", err)
	}
	fmt.Println("new uploadfile download", dow)
}

func Deleteobject(res http.ResponseWriter, req *http.Request) {
	var newfile file
	params := req.URL.Query()
	newfile.Bucketname=params["bucketname"][0]
	newfile.Objectname=params["objectname"][0]
	fmt.Fprintln(res,newfile)
	accessid := os.Getenv("s3accessid")
	accesskey := os.Getenv("secretacceskey")
	token := ""
	cred := credentials.NewStaticCredentials(accessid, accesskey, token)
	config := aws.NewConfig()
	config.Credentials = cred
	config.Region = aws.String("us-east-2")
	svc := s3.New((session.New(config)))
	buckname := newfile.Bucketname
	objname := newfile.Objectname
	deleteout, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(buckname), Key: aws.String(objname)})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res,deleteout)
}