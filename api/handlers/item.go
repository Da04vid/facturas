package handlers

import (
	"encoding/json"
	"facturas/models"
	"facturas/repository"
	"facturas/server"
	"net/http"
)

type ItemRequest struct {
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Valor string `json:"valor"`
}

type BuscarItemRequest struct {
	Id_item int64 `json:"id_item"`
}

type ItemResponse struct {
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Valor string `json:"valor"`
}

func ItemHandler(s server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var request = ItemRequest{}
		err:=json.NewDecoder(r.Body).Decode(&request) 
		if err != nil{
			http.Error(w, err.Error(),http.StatusBadRequest)
			return 
		}
		var item = models.Item{
			Nombre: request.Nombre,
			Descripcion: request.Descripcion,
			Valor: request.Valor,
		}
		err = repository.InsertItem(r.Context(),&item)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(ItemResponse{
			Nombre: item.Nombre,
			Descripcion: item.Descripcion,
			Valor: item.Valor,
		})
	}
}

func GetItemIdHandler(s server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var request = BuscarItemRequest{}
		err:=json.NewDecoder(r.Body).Decode(&request) 
		if err != nil{
			http.Error(w, err.Error(),http.StatusBadRequest)
			return 
		}
		item, err := repository.GetItemById(r.Context(),request.Id_item)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(ItemResponse{
			Nombre: item.Nombre,
			Descripcion: item.Descripcion,
			Valor: item.Valor,
		})
	}	
}

func GetItemsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := repository.GetItems(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
