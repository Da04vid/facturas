package database

import (
	"context"
	"facturas/models"
	"log"
	_"github.com/lib/pq"
)

func (repo *PostgresRepository) InsertFactura(ctx context.Context, factura *models.Factura) error{
	_, err := repo.db.ExecContext(ctx,"INSERT INTO factura (id_factura,fecha,descripcion,id_cliente) VALUES ($1,$2,$3,$4)",
	factura.Id_factura,factura.Fecha,factura.Descripcion,factura.Id_cliente)
	return err
}

func (repo *PostgresRepository) GetFacturaById(ctx context.Context, id string) (*models.Factura, error){
	filas, err := repo.db.QueryContext(ctx,"SELECT id_factura,fecha,descripcion,id_cliente FROM factura WHERE id_factura = $1",id)
	
	defer func(){
		err = filas.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()
	var factura = models.Factura{}
	for filas.Next(){
		if err = filas.Scan(&factura.Id_factura,&factura.Fecha,&factura.Descripcion,&factura.Id_cliente); err ==  nil{
			return &factura,nil
		}
	}
	if err = filas.Err();err != nil{
		return nil,err
	}
	return &factura,nil
}

func (repo *PostgresRepository) GetFacturas(ctx context.Context) ([]*models.Factura, error){
	filas, err := repo.db.QueryContext(ctx,"SELECT id_factura, descripcion FROM factura")
	
	defer func(){
		err = filas.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()
	var facturas []*models.Factura
	for filas.Next() {
		var factura models.Factura
		if err := filas.Scan(&factura.Id_factura,&factura.Fecha,&factura.Descripcion,&factura.Id_cliente); err != nil {
			return nil, err
		}
		facturas = append(facturas, &factura)
	}
	if err := filas.Err(); err != nil {
		return nil, err
	}

	return facturas, nil
}
