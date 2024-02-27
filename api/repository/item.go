package repository

import (
	"context"
	"facturas/models"
)

type RepositoryItem interface {
	InsertItem(ctx context.Context, Item *models.Item) error
	GetItemById(ctx context.Context, id int64) (*models.Item, error)
	GetItems(ctx context.Context) ([]*models.Item, error)
}

var implementationItem RepositoryItem

func SetRepositoryItem(repository RepositoryItem){
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





