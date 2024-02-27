package models

type FacturaItem struct {
	Id_factura_item     int64  `json:"id_factura_item"`
	Id_factura          int64  `json:"id_factura"`
	Id_item       int64 `json:"id_item"`
}
