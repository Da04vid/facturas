package server

import (
	"context"
	"errors"
	"facturas/database"
	"facturas/repository"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config{
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error){
	if config.Port ==  ""{
		return nil, errors.New("Puerto requerido")
	}
	if config.JWTSecret ==  ""{
		return nil, errors.New("Clave secreta requerida")
	}
	if config.DatabaseUrl ==  ""{
		return nil, errors.New("Base de datos requerida")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server,r *mux.Router)){
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil{
		log.Fatal(err)
	}
	repository.SetRepositoryCliente(repo)
	repository.SetRepositoryFactura(repo)
	repository.SetRepositoryItem(repo)
	log.Print("Inicializando servidor", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
