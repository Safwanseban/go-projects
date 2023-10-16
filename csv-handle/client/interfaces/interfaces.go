package interfaces

import "io"

type AwsClientI interface {
	UploadToS3(bucket, key string, data io.ReadSeeker) error
}
