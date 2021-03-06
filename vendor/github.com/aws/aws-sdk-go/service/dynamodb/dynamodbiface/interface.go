// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package dynamodbiface provides an interface for the Amazon DynamoDB.
package dynamodbiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBAPI is the interface type for dynamodb.DynamoDB.
type DynamoDBAPI interface {
	BatchGetItemRequest(*dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput)

	BatchGetItem(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)

	BatchGetItemPages(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error

	BatchWriteItemRequest(*dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput)

	BatchWriteItem(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)

	CreateTableRequest(*dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput)

	CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)

	DeleteItemRequest(*dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput)

	DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)

	DeleteTableRequest(*dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput)

	DeleteTable(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error)

	DescribeTableRequest(*dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput)

	DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)

	GetItemRequest(*dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput)

	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)

	ListTablesRequest(*dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput)

	ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)

	ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error

	PutItemRequest(*dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput)

	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)

	QueryRequest(*dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput)

	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)

	QueryPages(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error

	ScanRequest(*dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput)

	Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)

	ScanPages(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error

	UpdateItemRequest(*dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput)

	UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)

	UpdateTableRequest(*dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput)

	UpdateTable(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error)
}

var _ DynamoDBAPI = (*dynamodb.DynamoDB)(nil)
