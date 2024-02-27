package models

type Item struct{
	Id_item int64 `json:"id_item"`
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Valor string `json:"valor"`
}
