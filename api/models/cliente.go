package models

type Cliente struct{
	Id_cliente int64 `json:"id"`
	Nombre string `json:"nombre"`
	Telefono string `json:"telefono"`
	Identificacion string `json:"identificacion"`
	Correo string `json:"correo"`
}
