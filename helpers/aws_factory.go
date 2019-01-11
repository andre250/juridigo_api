package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juridigo/juridigo_api_usuario/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var preserveDirStructureBool bool
var awsSession *session.Session

type AWSession struct {
	Session *session.Session
}

/*
MakeSession - Criar sessão Amazon
*/
func MakeSession() {
	configuration = config.GetConfig()
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(configuration.Amazon.Region)},
	})
	if err != nil {
		fmt.Println("failed to create session,", err)
		fmt.Println(err)
		os.Exit(1)
	}
	awsSession = sess
}

/*
AWS - Obter sessão amazon
*/
func AWS() *AWSession {
	session := AWSession{
		Session: awsSession,
	}
	return &session
}

/*
UploadFileToS3 - Executa o upload do arquivo
*/
func (s *AWSession) UploadFileToS3(path string) {
	s3Svc := s3.New(s.Session)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open file", file, err)
		os.Exit(1)
	}
	defer file.Close()
	var key string
	if preserveDirStructureBool {
		fileDirectory, _ := filepath.Abs(path)
		key = configuration.Amazon.Prefix + fileDirectory
	} else {
		key = configuration.Amazon.Prefix + path
	}
	// Upload the file to the s3 given bucket
	params := &s3.PutObjectInput{
		Bucket: aws.String(configuration.Amazon.Bucket), // Required
		Key:    aws.String(key),                         // Required
		Body:   file,
	}
	_, err = s3Svc.PutObject(params)
	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n",
			"juridigo", key, err.Error())
		return
	}
}
