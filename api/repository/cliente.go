package repository

import (
	"context"
	"facturas/models"
)

type ClienteRepository interface {
	InsertCliente(ctx context.Context, cliente *models.Cliente) error
	GetClienteByIdentificacion(ctx context.Context, identificacion string) (*models.Cliente, error)
	GetClientes(ctx context.Context) ([]*models.Cliente, error)
	Close() error
}

var implementation ClienteRepository

func SetRepositoryCliente(repository ClienteRepository){
	implementation = repository
}

func InsertCliente(ctx context.Context, cliente *models.Cliente) error{
	return implementation.InsertCliente(ctx,cliente)
}

func GetClienteByIdentificacion(ctx context.Context, identificaion string) (*models.Cliente,error){
	return implementation.GetClienteByIdentificacion(ctx,identificaion)
}

func GetClientes(ctx context.Context) ([]*models.Cliente,error){
	return implementation.GetClientes(ctx)
}

func Close() error{
	return implementation.Close()
}


