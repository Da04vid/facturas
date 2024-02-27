package models

type Item struct{
	Id_item int64 `json:"id"`
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Valor string `json:"valor"`
}
