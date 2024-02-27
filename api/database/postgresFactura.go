package database

import (
	"context"
	"facturas/models"
	"log"
	_"github.com/lib/pq"
)

func (repo *PostgresRepository) InsertFactura(ctx context.Context, cliente *models.Cliente) error{
	_, err := repo.db.ExecContext(ctx,"INSERT INTO factura (nombre,telefono,identificacion,correo) VALUES ($1,$2,$3,$4)",
	cliente.Nombre,cliente.Telefono,cliente.Identificacion,cliente.Correo)
	return err
}

func (repo *PostgresRepository) GetFacturaById(ctx context.Context, identificacion string) (*models.Cliente, error){
	filas, err := repo.db.QueryContext(ctx,"SELECT nombre,telefono,identificacion,correo FROM factura WHERE id = $1",identificacion)
	
	defer func(){
		err = filas.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()
	var cliente = models.Cliente{}
	for filas.Next(){
		if err = filas.Scan(&cliente.Nombre, &cliente.Telefono, &cliente.Identificacion, &cliente.Correo); err ==  nil{
			return &cliente,nil
		}
	}
	if err = filas.Err();err != nil{
		return nil,err
	}
	return &cliente,nil
}

func (repo *PostgresRepository) GetFacturas(ctx context.Context) ([]*models.Cliente, error){
	filas, err := repo.db.QueryContext(ctx,"SELECT id_factura, descripcion FROM factura")
	
	defer func(){
		err = filas.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()
	var clientes []*models.Cliente
	for filas.Next() {
		var cliente models.Cliente
		if err := filas.Scan(&cliente.Nombre, &cliente.Telefono, &cliente.Identificacion, &cliente.Correo); err != nil {
			return nil, err
		}
		clientes = append(clientes, &cliente)
	}
	if err := filas.Err(); err != nil {
		return nil, err
	}

	return clientes, nil
}
