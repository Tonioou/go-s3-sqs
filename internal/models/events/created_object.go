package models

type Message struct {
	Records []Records `json:"Records"`
}

type Records struct {
	S3 S3 `json:"s3"`
}

type S3 struct {
	Object Object `json:"object"`
}

type Object struct {
	Key string `json:"key"`
}
