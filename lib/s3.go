package lib

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// PresignedURLExpiration is the duration for Presigned URL to expire
const PresignedURLExpiration = 10 * time.Minute

// S3PresignedURL is the return type for PutS3PresignedURL method
type S3PresignedURL struct {
	URL       string    `json:"url"`
	ExpiresAt time.Time `json:"expires_at"`
}

// GetS3PresignedURL is used to get a presigned url to PUT an asset
func GetS3PresignedURL(key string) (S3PresignedURL, error) {
	// Initialize a session in the target region that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.

	// Create S3 service client
	svc := s3.New(session.New())

	// Construct a GetObjectRequest request
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("dms_bucket")),
		Key:    aws.String(key),
	})

	var psURL S3PresignedURL
	// Presign with expiration time
	url, err := req.Presign(PresignedURLExpiration)
	if err != nil {
		return psURL, err
	}

	psURL = S3PresignedURL{
		URL:       url,
		ExpiresAt: time.Now().Add(PresignedURLExpiration),
	}

	// Return the presigned url
	return psURL, nil
}

// PutS3PresignedURL is used to get a presigned url to PUT an asset
func PutS3PresignedURL(key string) (S3PresignedURL, error) {
	// Initialize a session in the target region that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.

	// Create S3 service client
	svc := s3.New(session.New())

	// Construct a GetObjectRequest request
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("dms_bucket")),
		Key:    aws.String(key),
	})

	var psURL S3PresignedURL
	// Presign with expiration time
	url, err := req.Presign(PresignedURLExpiration)
	if err != nil {
		return psURL, err
	}

	psURL = S3PresignedURL{
		URL:       url,
		ExpiresAt: time.Now().Add(PresignedURLExpiration),
	}

	// Return the presigned url
	return psURL, nil
}
