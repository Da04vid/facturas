package repository

import (
	"context"
	"facturas/models"
)

type FacturaRepository interface {
	InsertFactura(ctx context.Context, factura *models.Factura) error
	GetFacturaById(ctx context.Context, id string) (*models.Factura, error)
	GetFacturas(ctx context.Context) ([]*models.Factura, error)
}

var implementationFactura FacturaRepository

func SetRepositoryFactura(repository FacturaRepository){
	implementationFactura = repository
}

func InsertFactura(ctx context.Context, factura *models.Factura) error{
	return implementationFactura.InsertFactura(ctx,factura)
}

func GetFacturaById(ctx context.Context, id string) (*models.Factura,error){
	return implementationFactura.GetFacturaById(ctx,id)
}

func GetFacturas(ctx context.Context) ([]*models.Factura,error){
	return implementationFactura.GetFacturas(ctx)
}


