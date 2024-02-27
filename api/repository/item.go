package repository

import (
	"context"
	"facturas/models"
)

type ItemRepository interface {
	InsertItem(ctx context.Context, Item *models.Item) error
	GetItemById(ctx context.Context, id int64) (*models.Item, error)
	GetItems(ctx context.Context) ([]*models.Item, error)
}

var implementationItem ItemRepository

func SetRepositoryItem(repository ItemRepository){
	implementationItem = repository
}

func InsertItem(ctx context.Context, item *models.Item) error{
	return implementationItem.InsertItem(ctx,item)
}

func GetItemById(ctx context.Context, id int64) (*models.Item,error){
	return implementationItem.GetItemById(ctx,id)
}

func GetItems(ctx context.Context) ([]*models.Item,error){
	return implementationItem.GetItems(ctx)
}

