package repositories

import (
	models "mrkresnofatih/gobookapi/models"

	"log"

	aws "github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	dynamodbattribute "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getDynamoDbSvc() (*dynamodb.DynamoDB, error) {
	// ENV AWS_ACCESS_KEY_ID & AWS_SECRET_ACCESS_KEY
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	svc := dynamodb.New(sess)
	return svc, nil
}

func AddBook(book models.Book) (*models.Book, error) {
	svc, errr := getDynamoDbSvc()
	if errr != nil {
		log.Println(errr)
		return nil, errr
	}

	info, err := dynamodbattribute.MarshalMap(book)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	putRequest := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("GoBookAPI-Books"),
	}

	_, er1 := svc.PutItem(putRequest)
	if er1 != nil {
		log.Println(er1)
		return nil, er1
	}

	return &book, nil
}

func GetBook(id string) (*models.Book, error) {
	svc, err := getDynamoDbSvc()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(id)
	key := map[string]*dynamodb.AttributeValue{
		"bookId": {
			S: aws.String(id),
		},
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String("GoBookAPI-Books"),
	}
	result, er := svc.GetItem(input)
	if er != nil {
		log.Println(er)
		return nil, er
	}

	book := models.Book{}
	errr := dynamodbattribute.UnmarshalMap(result.Item, &book)
	if errr != nil {
		log.Println(errr)
		return nil, errr
	}
	return &book, nil
}
