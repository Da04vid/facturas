package handlers

import (
	"encoding/json"
	"facturas/models"
	"facturas/repository"
	"facturas/server"
	"net/http"
)

type ClienteRequest struct {
	Nombre string `json:"nombre"`
}

type ClienteResponse struct {
	Nombre         string `json:"nombre"`
	Telefono       string `json:"telefono"`
	Identificacion string `json:"identificacion"`
	Correo         string `json:"correo"`
}

func ClienteHandler(s server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var request = ClienteRequest{}
		err:=json.NewDecoder(r.Body).Decode(&request) 
		if err != nil{
			http.Error(w, err.Error(),http.StatusBadRequest)
			return 
		}
		var cliente = models.Cliente{
			Nombre: request.Nombre,
		}
		err = repository.InsertCliente(r.Context(),&cliente)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(ClienteResponse{
			Nombre: cliente.Nombre,
			Telefono: cliente.Telefono,
			Identificacion: cliente.Identificacion,
			Correo: cliente.Correo,
		})
	}
}
