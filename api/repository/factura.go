package repository

import (
	"context"
	"facturas/models"
)

type RepositoryFactura interface {
	InsertFactura(ctx context.Context, factura *models.Factura) error
	GetFacturaById(ctx context.Context, id int64) (*models.Factura, error)
	GetFacturas(ctx context.Context) ([]*models.Factura, error)
}

var implementationFactura RepositoryFactura

func SetRepositoryFactura(repository RepositoryFactura){
	implementationFactura = repository
}

func InsertFactura(ctx context.Context, factura *models.Factura) error{
	return implementationFactura.InsertFactura(ctx,factura)
}

func GetFacturaById(ctx context.Context, id int64) (*models.Factura,error){
	return implementationFactura.GetFacturaById(ctx,id)
}

func GetFacturas(ctx context.Context) ([]*models.Factura,error){
	return implementationFactura.GetFacturas(ctx)
}


