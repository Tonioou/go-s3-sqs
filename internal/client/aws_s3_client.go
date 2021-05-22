package client

type S3Bucket interface {
	GetObject(path string)
}
