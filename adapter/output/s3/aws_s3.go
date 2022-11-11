package s3

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateSession() *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("REGION")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("ACCESS_KEY_ID"),
				os.Getenv("SECRET_ACCESS_KEY"),
				"",
			),
		},
	))
	return sess
}

func UploadObject(id string, docType string, fileName string, fileBytes []byte) (string, error) {
	session := CreateSession()

	uploader := s3manager.NewUploader(session)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("odontooth"),
		Key:    aws.String(fmt.Sprintf("%s/%s/%s.jpeg", id, docType, fileName)),
		Body:   bytes.NewReader(fileBytes),
	})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return result.Location, nil
}
