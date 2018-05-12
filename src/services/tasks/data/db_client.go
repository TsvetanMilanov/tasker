package data

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/TsvetanMilanov/tasker-common/common"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DBClient is the bridge to the database.
type DBClient struct {
	clients map[string]*dynamodb.DynamoDB
}

// Create saves the provided item to the provided table.
func (c *DBClient) Create(item interface{}, table string, awsCfg *aws.Config) error {
	err := c.validateInput(item)
	if err != nil {
		return err
	}

	i, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	if _, ok := i[createdOnKey]; !ok {
		att, err := dynamodbattribute.ConvertTo(time.Now().UTC())
		if err != nil {
			return err
		}

		i[createdOnKey] = att
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      i,
	}
	client := c.getDynamoDBClient(table, awsCfg)
	_, err = client.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (c *DBClient) getDynamoDBClient(table string, awsCfg *aws.Config) *dynamodb.DynamoDB {
	if c.clients[table] == nil {
		sess := session.Must(session.NewSession(awsCfg))
		c.clients[table] = dynamodb.New(sess)
	}

	return c.clients[table]
}

func (c *DBClient) validateInput(item interface{}) error {
	if item == nil {
		return errors.New("the provided item is nil")
	}

	kind := reflect.TypeOf(item).Kind()
	if kind != reflect.Struct {
		return errors.New("the provided item is not struct")
	}

	vErrs := common.ValidateStruct(item)

	var err error

	if len(vErrs) > 0 {
		messages := []string{}
		for _, e := range vErrs {
			messages = append(messages, e.Translate(nil))
		}

		err = errors.New(strings.Join(messages, "\n"))
	}

	return err
}

// NewDBClient returns new database client.
func NewDBClient() *DBClient {
	return &DBClient{clients: make(map[string]*dynamodb.DynamoDB)}
}
