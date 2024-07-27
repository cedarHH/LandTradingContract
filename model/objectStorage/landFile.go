package objectStorage

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type LandModel struct {
	*S3Model
}

func NewLandModel(client *s3.Client, presignClient *s3.PresignClient, bucket string) IS3Model {
	return &LandModel{
		&S3Model{
			Client:        client,
			PresignClient: presignClient,
			Bucket:        bucket,
		},
	}
}
