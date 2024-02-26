package main

import (
	"context"
	"facturas/handlers"
	"facturas/server"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error cargando")
	}
	PORT := os.Getenv("PORT")
	JWT_SECRET:=os.Getenv("JWT_SECRET")
	DATABASE_URL:=os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
		JWTSecret :JWT_SECRET,
		DatabaseUrl:DATABASE_URL,
	})

	if err != nil{
		log.Fatal(err)
	}
	s.Start(BinRoutes)
}

func BinRoutes(s server.Server, r *mux.Router){
	r.HandleFunc("/",handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/cliente",handlers.ClienteHandler(s)).Methods(http.MethodPost)
}