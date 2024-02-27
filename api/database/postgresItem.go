package database

import (
	"context"
	"facturas/models"
	"log"
	_"github.com/lib/pq"
)

func (repo *PostgresRepository) InsertItem(ctx context.Context, item *models.Item) error{
	_, err := repo.db.ExecContext(ctx,"INSERT INTO item (id_item,nombre,descripcion,valor) VALUES ($1,$2,$3,$4)",
	item.Id_item,item.Nombre,item.Descripcion,item.Valor)
	return err
}

func (repo *PostgresRepository) GetItemById(ctx context.Context, id string) (*models.Item, error){
	filas, err := repo.db.QueryContext(ctx,"SELECT id_item,nombre,descripcion,valor FROM item WHERE id_item = $1",id)
	
	defer func(){
		err = filas.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()
	var item = models.Item{}
	for filas.Next(){
		if err = filas.Scan(&item.Id_item,&item.Nombre,&item.Descripcion,&item.Valor); err ==  nil{
			return &item,nil
		}
	}
	if err = filas.Err();err != nil{
		return nil,err
	}
	return &item,nil
}

func (repo *PostgresRepository) GetItems(ctx context.Context) ([]*models.Item, error){
	filas, err := repo.db.QueryContext(ctx,"SELECT  id_item,nombre,descripcion,valor FROM item")
	
	defer func(){
		err = filas.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()
	var items []*models.Item
	for filas.Next() {
		var item models.Item
		if err := filas.Scan(&item.Id_item,&item.Nombre,&item.Descripcion,&item.Valor); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	if err := filas.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
