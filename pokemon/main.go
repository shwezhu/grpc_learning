package main

import (
	"github.com/shwezhu/grpc_learning/pokemon/database"
	"github.com/shwezhu/grpc_learning/pokemon/pb"
	"github.com/shwezhu/grpc_learning/pokemon/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	log.Println("Server is running on port :8080")

	db, err := database.Connect()
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	newServer := server.NewServer(server.NewRepository(db))
	s := grpc.NewServer()
	reflection.Register(s) // used for grpcurl
	pb.RegisterPokemonServer(s, newServer)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
