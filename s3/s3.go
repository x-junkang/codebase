// Amazon S3服务

package s3

import (
	"bytes"
	"codebase/utils"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const PartSize = 1 * 1024 * 1024
const MaxBodySize = 10 * 1024 * 1024

var Svc *s3.S3
var Bucket string

func S3Init(region string, accessKey string, secretKey string, endpoint string, bucket string) {
	sess := session.New(&aws.Config{
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})
	Svc = s3.New(sess)
	Bucket = bucket
}

func PutObject(key string, data io.ReadSeeker) error {
	_, err := Svc.PutObject(
		&s3.PutObjectInput{
			Body:   data,
			Key:    aws.String(key),
			Bucket: aws.String(Bucket),
		})

	if err != nil {
		return err
	}
	return nil
}

func PutObjectWithTagging(key string, tags map[string]string, data io.ReadSeeker) error {
	tagging := utils.MapToKVString(tags)
	input := &s3.PutObjectInput{
		Body:   data,
		Key:    aws.String(key),
		Bucket: aws.String(Bucket),
	}
	if tagging != "" {
		input.Tagging = aws.String(tagging)
	}
	if _, err := Svc.PutObject(input); err != nil {
		return err
	}
	return nil
}

func GetObject(key string) (int64, io.ReadCloser) {
	out, err := Svc.GetObject(
		&s3.GetObjectInput{
			Bucket: &Bucket,
			Key:    &key,
		},
	)

	if err != nil {
		return 0, nil
	}

	return *(out.ContentLength), out.Body
}

func PutMultiPartObject(key string, file *os.File) error {
	var err error
	var ioerr error

	var offset int64 = 0

	var i int64 = 1
	var partNum []int64
	var etag string
	var etagSlice []string
	var readBytes int
	var multiPart []*s3.CompletedPart

	buffer := make([]byte, PartSize)

	uploadId, err := createMultiPartUpload(key)
	if err != nil {
		return err
	}

	for {
		readBytes, ioerr = file.ReadAt(buffer, offset)
		if readBytes > 0 {
			etag, err = uploadPart(key, uploadId, i, bytes.NewReader(buffer[:readBytes]))
			if err != nil {
				return abortMultiPartUpload(key, uploadId)
			}
			partNum = append(partNum, i)
			etagSlice = append(etagSlice, etag)
			multiPart = append(multiPart, &s3.CompletedPart{PartNumber: &partNum[i-1], ETag: &etagSlice[i-1]})
			i++
			offset = (i - 1) * PartSize
		}
		if ioerr == io.EOF {
			return completeMultiPartUpload(key, uploadId, multiPart)
		}
		if ioerr != nil {
			return abortMultiPartUpload(key, uploadId)
		}

	}

}

func createMultiPartUpload(key string) (string, error) {
	input := &s3.CreateMultipartUploadInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(key),
	}

	output, err := Svc.CreateMultipartUpload(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return "", aerr
			}
		} else {

			return "", err
		}
	}

	return *output.UploadId, nil

}

func PutMultiPartObjectWithTagging(key string, tags map[string]string, file *os.File) error {
	var (
		err       error
		ioerr     error
		offset    int64 = 0
		i         int64 = 1
		partNum   []int64
		etag      string
		etagSlice []string
		readBytes int
		multiPart []*s3.CompletedPart
	)
	buffer := make([]byte, PartSize)

	uploadId, err := createMultiPartUploadWithTagging(key, tags)
	if err != nil {
		return err
	}

	for {
		readBytes, ioerr = file.ReadAt(buffer, offset)
		if readBytes > 0 {
			etag, err = uploadPart(key, uploadId, i, bytes.NewReader(buffer[:readBytes]))
			if err != nil {
				return abortMultiPartUpload(key, uploadId)
			}
			partNum = append(partNum, i)
			etagSlice = append(etagSlice, etag)
			multiPart = append(multiPart, &s3.CompletedPart{PartNumber: &partNum[i-1], ETag: &etagSlice[i-1]})
			i++
			offset = (i - 1) * PartSize
		}
		if ioerr == io.EOF {
			return completeMultiPartUpload(key, uploadId, multiPart)
		}
		if ioerr != nil {
			return abortMultiPartUpload(key, uploadId)
		}
	}
}

func createMultiPartUploadWithTagging(key string, tags map[string]string) (string, error) {
	tagging := utils.MapToKVString(tags)
	input := &s3.CreateMultipartUploadInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(key),
	}
	if tagging != "" {
		input.Tagging = aws.String(tagging)
	}

	output, err := Svc.CreateMultipartUpload(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return "", aerr
			}
		} else {

			return "", err
		}
	}

	return *output.UploadId, nil

}

func uploadPart(key string, uploadId string, partNum int64, buffer io.ReadSeeker) (string, error) {
	input := &s3.UploadPartInput{
		Body:       aws.ReadSeekCloser(buffer),
		Bucket:     aws.String(Bucket),
		Key:        aws.String(key),
		PartNumber: aws.Int64(partNum),
		UploadId:   aws.String(uploadId),
	}

	output, err := Svc.UploadPart(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return "", aerr
			}
		} else {
			return "", err
		}
	}
	return *output.ETag, nil
}

func abortMultiPartUpload(key string, uploadId string) error {
	input := &s3.AbortMultipartUploadInput{
		Bucket:   aws.String(Bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadId),
	}

	_, err := Svc.AbortMultipartUpload(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return aerr
			}
		} else {
			return err
		}
	}
	return nil
}

func completeMultiPartUpload(key string, uploadId string, parts []*s3.CompletedPart) error {
	input := &s3.CompleteMultipartUploadInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(key),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: parts,
		},
		UploadId: aws.String(uploadId),
	}

	_, err := Svc.CompleteMultipartUpload(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return aerr
			}
		} else {
			return err
		}
	}
	return nil
}
