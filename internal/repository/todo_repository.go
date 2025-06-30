package repository

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"echo-todo/pkg/models"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *models.Todo) error
	GetByID(ctx context.Context, id string) (*models.Todo, error)
	GetAll(ctx context.Context) ([]models.Todo, error)
	Update(ctx context.Context, todo *models.Todo) error
	Delete(ctx context.Context, id string) error
}

type DynamoDBTodoRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDBTodoRepository(tableName string) (*DynamoDBTodoRepository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return &DynamoDBTodoRepository{
		client:    client,
		tableName: tableName,
	}, nil
}

func (r *DynamoDBTodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	item, err := attributevalue.MarshalMap(todo)
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	})
	return err
}

func (r *DynamoDBTodoRepository) GetByID(ctx context.Context, id string) (*models.Todo, error) {
	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
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

	var todo models.Todo
	err = attributevalue.UnmarshalMap(result.Item, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *DynamoDBTodoRepository) GetAll(ctx context.Context) ([]models.Todo, error) {
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	})
	if err != nil {
		return nil, err
	}

	var todos []models.Todo
	err = attributevalue.UnmarshalListOfMaps(result.Items, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *DynamoDBTodoRepository) Update(ctx context.Context, todo *models.Todo) error {
	item, err := attributevalue.MarshalMap(todo)
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	})
	return err
}

func (r *DynamoDBTodoRepository) Delete(ctx context.Context, id string) error {
	_, err := r.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}