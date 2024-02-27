package handlers

import (
	"encoding/json"
	"facturas/models"
	"facturas/repository"
	"facturas/server"
	"net/http"
)

type ClienteRequest struct {
	Nombre         string `json:"nombre"`
	Telefono       string `json:"telefono"`
	Identificacion string `json:"identificacion"`
	Correo         string `json:"correo"`
}

type BuscarClienteRequest struct {
	Identificacion string `json:"identificacion"`
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
			Telefono: request.Telefono,
			Correo: request.Correo,
			Identificacion: request.Identificacion,
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

func GetIdentificacionHandler(s server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var request = BuscarClienteRequest{}
		err:=json.NewDecoder(r.Body).Decode(&request) 
		if err != nil{
			http.Error(w, err.Error(),http.StatusBadRequest)
			return 
		}
		cliente, err := repository.GetClienteByIdentificacion(r.Context(),request.Identificacion)
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

func GetClientesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientes, err := repository.GetClientes(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(clientes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
