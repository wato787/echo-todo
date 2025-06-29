package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DatabaseClient struct {
	DynamoDB  *dynamodb.Client
	TableName string
}

func NewDatabaseClient(tableName string) (*DatabaseClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return &DatabaseClient{
		DynamoDB:  client,
		TableName: tableName,
	}, nil
}

func (db *DatabaseClient) CreateTodo(ctx context.Context, todo *Todo) error {
	item, err := attributevalue.MarshalMap(todo)
	if err != nil {
		return err
	}

	_, err = db.DynamoDB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.TableName),
		Item:      item,
	})
	return err
}

func (db *DatabaseClient) GetTodo(ctx context.Context, id string) (*Todo, error) {
	result, err := db.DynamoDB.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var todo Todo
	err = attributevalue.UnmarshalMap(result.Item, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (db *DatabaseClient) GetAllTodos(ctx context.Context) ([]Todo, error) {
	result, err := db.DynamoDB.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(db.TableName),
	})
	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = attributevalue.UnmarshalListOfMaps(result.Items, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (db *DatabaseClient) UpdateTodo(ctx context.Context, todo *Todo) error {
	item, err := attributevalue.MarshalMap(todo)
	if err != nil {
		return err
	}

	_, err = db.DynamoDB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.TableName),
		Item:      item,
	})
	return err
}

func (db *DatabaseClient) DeleteTodo(ctx context.Context, id string) error {
	_, err := db.DynamoDB.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(db.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}