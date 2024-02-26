package repository

import (
	"context"
	"facturas/models"
)

type ClienteRepository interface {
	InsertCliente(ctx context.Context, cliente *models.Cliente) error
	GetClienteById(ctx context.Context, id int64) (*models.Cliente, error)
}

var implementation ClienteRepository

func SetRepository(repository ClienteRepository){
	implementation = repository
}

func InsertCliente(ctx context.Context, cliente *models.Cliente) error{
	return implementation.InsertCliente(ctx,cliente)
}

func GetClienteById(ctx context.Context, id int64) (*models.Cliente,error){
	return implementation.GetClienteById(ctx,id)
}
