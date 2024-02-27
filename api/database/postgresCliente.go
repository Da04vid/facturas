package database

import (
	"context"
	"database/sql"
	"facturas/models"
	_ "github.com/lib/pq"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertCliente(ctx context.Context, cliente *models.Cliente) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO cliente (nombre,telefono,identificacion,correo) VALUES ($1,$2,$3,$4)",
		cliente.Nombre, cliente.Telefono, cliente.Identificacion, cliente.Correo)
	return err
}

func (repo *PostgresRepository) GetClienteByIdentificacion(ctx context.Context, identificacion string) (*models.Cliente, error) {
	filas, err := repo.db.QueryContext(ctx, "SELECT nombre,telefono,identificacion,correo FROM cliente WHERE identificacion = $1", identificacion)

	defer func() {
		err = filas.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var cliente = models.Cliente{}
	for filas.Next() {
		if err = filas.Scan(&cliente.Nombre, &cliente.Telefono, &cliente.Identificacion, &cliente.Correo); err == nil {
			return &cliente, nil
		}
	}
	if err = filas.Err(); err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (repo *PostgresRepository) GetClientes(ctx context.Context) ([]*models.Cliente, error) {
	filas, err := repo.db.QueryContext(ctx, "SELECT nombre,telefono,identificacion,correo FROM cliente")

	defer func() {
		err = filas.Close()
		if err != nil {
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

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
