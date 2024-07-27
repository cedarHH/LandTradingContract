package database

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Land represents the land entity stored in DynamoDB
type Land struct {
	LandId   string `json:"landID" dynamodbav:"landId"`
	Detail   string `json:"Detail" dynamodbav:"detail"`
	Report   string `json:"Report" dynamodbav:"report"`
	Document string `json:"Document" dynamodbav:"document"`
}

type LandModel interface {
	Insert(ctx context.Context, land *Land) error
	FindOne(ctx context.Context, landId string) (*Land, error)
	Update(ctx context.Context, land *Land) error
	Delete(ctx context.Context, landId string) error
	UpdateAttributes(ctx context.Context, landId string, updates map[string]interface{}) error
}

type landModel struct {
	baseModel *BaseModel
}

func NewLandModel(client *dynamodb.Client, table string) LandModel {
	return &landModel{
		&BaseModel{
			Client: client,
			Table:  table,
		},
	}
}

func (m *landModel) Insert(ctx context.Context, land *Land) error {
	return m.baseModel.Insert(ctx, land)
}

func (m *landModel) FindOne(ctx context.Context, landId string) (*Land, error) {
	key := map[string]interface{}{
		"landId": landId,
	}

	item, err := m.baseModel.FindOne(ctx, key)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}

	var land Land
	err = attributevalue.UnmarshalMap(item, &land)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal land: %w", err)
	}
	return &land, nil
}

func (m *landModel) Update(ctx context.Context, land *Land) error {
	return m.baseModel.Update(ctx, land)
}

func (m *landModel) Delete(ctx context.Context, landId string) error {
	key := map[string]interface{}{
		"landId": landId,
	}
	return m.baseModel.Delete(ctx, key)
}

func (m *landModel) UpdateAttributes(ctx context.Context, landId string, updates map[string]interface{}) error {
	key := map[string]interface{}{
		"landId": landId,
	}
	return m.baseModel.UpdateAttributes(ctx, key, updates)
}
