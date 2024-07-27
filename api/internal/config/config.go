package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DynamoDBConf DynamoDBConf
	S3Conf       S3Conf
}

type DynamoDBConf struct {
	Region    string `json:"Region"`
	TableName string `json:"TableName"`
}

type S3Conf struct {
	Region string `json:"Region"`
	Bucket string `json:"Bucket"`
}
