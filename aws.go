/*
*
* Author: Hui Ye - <bonjovis@163.com>
*
* Last modified: 2017-03-21 08:21
*
* Filename: aws.go
*
* Copyright (c) 2016 JOVI
*
 */
package awsservice

import "fmt"
import "bytes"
import "time"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/aws/credentials"

//import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/service/sqs"
import "github.com/aws/aws-sdk-go/service/s3"

//aws
var _ time.Duration
var _ bytes.Buffer

type AwsService struct {
	session *session.Session
}

func AwsSessionHardCredentials(regin, awsKey, secret string) *AwsService {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(regin),
		Credentials: credentials.NewStaticCredentials(awsKey, secret, ""),
	})
	if err != nil {
		panic(err)
	}
	aService := &AwsService{sess}
	return aService
}

func AwsSessionSharedCredentials() *AwsService {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	aService := &AwsService{sess}
	return aService
}

func (aService *AwsService) ListObjectKeys(bucket, prefix string) []string {
	svc := s3.New(aService.session)
	inputparams := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}
	resp, _ := svc.ListObjects(inputparams)
	list := make([]string, len(resp.Contents))
	for _, key := range resp.Contents {
		list = append(list, *key.Key)
	}
	return list
}

func (aService *AwsService) GetObject(bucket, key string) string {
	svc := s3.New(aService.session)
	getparams := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	resp, err := svc.GetObject(getparams)
	if err != nil {
		fmt.Println("The Key Failed ", key)
		fmt.Println("Failed to get object", err)
	}
	var s string
	if resp.Body != nil {
		s = StreamToString(resp.Body)
	}
	defer resp.Body.Close()
	return s
}

func (aService *AwsService) PutObject(bucket, key, body string) {
	svc := s3.New(aService.session)
	putparams := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader([]byte(body)),
	}
	_, err := svc.PutObject(putparams)
	if err != nil {
		fmt.Println("Failed to put object", err)
	}
}

func (aService *AwsService) SendMessageToSQS(qUrl string, qBody string, qAttribute map[string]*sqs.MessageAttributeValue) {
	svc := sqs.New(aService.session)

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:      aws.Int64(10),
		MessageAttributes: qAttribute,
		MessageBody:       aws.String(qBody),
		QueueUrl:          &qUrl,
	})

	if err != nil {
		fmt.Println("Failed to send message", err)
		return
	}
	fmt.Println("Success", *result.MessageId)
}

func (aService *AwsService) ReceiveMessageFromSQS(qUrl string) *sqs.ReceiveMessageOutput {
	svc := sqs.New(aService.session)

	params := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		QueueUrl:            aws.String(qUrl),
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(20),
	}
	resp, err := svc.ReceiveMessage(params)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Failed to receive message", err)
	}

	if len(resp.Messages) == 0 {
		fmt.Println("Received no messages")
	}
	return resp

}
func (aService *AwsService) DeleteMessage(msg *sqs.Message, qUrl string) {
	svc := sqs.New(aService.session)

	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(qUrl),
		ReceiptHandle: aws.String(*msg.ReceiptHandle),
	}
	_, err := svc.DeleteMessage(params)

	if err != nil {
		fmt.Println("Failed to delete message", err)
	}
}
