package models

type Factura struct {
	Id_factura     int64  `json:"id"`
	Fecha          string  `json:"fecha"`
	Descripcion       string `json:"descripcion"`
	Id_cliente int64 `json:"id_cliente"`
}
