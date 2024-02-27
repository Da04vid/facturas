package handlers

import (
	"encoding/json"
	"facturas/models"
	"facturas/repository"
	"facturas/server"
	"net/http"
)

type FacturasRequest struct {
	Fecha         string `json:"fecha"`
	Descripcion       string `json:"descripcion"`
	Id_cliente string `json:"id_cliente"`
}

type BuscarFacturaRequest struct {
	Id_factura int64 `json:"id_factura"`
}

type FacturaResponse struct {
	Fecha         string `json:"fecha"`
	Descripcion       string `json:"descripcion"`
	Id_cliente string `json:"id_cliente"`
}

func FacturaHandler(s server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var request = FacturasRequest{}
		err:=json.NewDecoder(r.Body).Decode(&request) 
		if err != nil{
			http.Error(w, err.Error(),http.StatusBadRequest)
			return 
		}
		var factura = models.Factura{
			Fecha: request.Fecha,
			Descripcion: request.Descripcion,
			Id_cliente: request.Id_cliente,
		}
		err = repository.InsertFactura(r.Context(),&factura)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(FacturaResponse{
			Fecha: factura.Fecha,
			Descripcion: factura.Descripcion,
			Id_cliente: factura.Id_cliente,
		})
	}
}

func GetFacturaIdHandler(s server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var request = BuscarFacturaRequest{}
		err:=json.NewDecoder(r.Body).Decode(&request) 
		if err != nil{
			http.Error(w, err.Error(),http.StatusBadRequest)
			return 
		}
		factura, err := repository.GetFacturaById(r.Context(),request.Id_factura)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(FacturaResponse{
			Fecha: factura.Fecha,
			Descripcion: factura.Descripcion,
			Id_cliente: factura.Id_cliente,
		})
	}	
}

func GetFacturasHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		facturas, err := repository.GetFacturas(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(facturas); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
