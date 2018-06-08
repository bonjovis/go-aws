# go-aws
Golang language integration of aws operations

## Table of contents:
- [Get Started](#get-started)
- [Examples](#examples)


### Get Started
#### Installation

```sh
$ go get github.com/aws/aws-sdk-go 
$ go get github.com/bonjovis/go-aws
```


#### Examples
```go
import (
	"fmt"
	"github.com/bonjovis/go-aws"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
    awsSession := awsservice.AwsSessionSharedCredentials()
    //s3
    awsSession.PutObject(bucket, filename, data)
    //ddb
    params := &dynamodb.PutItemInput{
        Item: map[string]*dynamodb.AttributeValue{
            "test1": {
                S: aws.String("t1"),
            },
            "test2": {
                N: aws.String("t2"),
            },
            "test3": {
                S: aws.String("t3"),
            },
        },
        TableName: aws.String("ddbTable"),
    }   
    awsSession.PutItem(params)
    //recive sqs message
    awsSession.ReceiveMessageFromSQS(sqsUrl, size)
    //send sqs message
    qBody :="test"
    qAttribute := map[string]*sqs.MessageAttributeValue{
                "test1": &sqs.MessageAttributeValue{
                    DataType:    aws.String("String"),
                    StringValue: aws.String("t1"),
                },
                "test2": &sqs.MessageAttributeValue{
                    DataType:    aws.String("String"),
                    StringValue: aws.String("t2"),
                },
                "test2": &sqs.MessageAttributeValue{
                    DataType:    aws.String("String"),
                    StringValue: aws.String("t3"),
                },
     }
     awsSession.SendMessageToSQS(sqsUrl, qBody, qAttribute)
     //delete message
     task.awsSession.DeleteMessage(mess, sqsUrl)
     //publish message
     awsSession.PublishMessageToSNS(object, snsUrl)
}
```


### License
MIT
