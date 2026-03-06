package app

import (
	"fmt"
	"log"
	"net"
	"os"
	filmsgrpc "service-2/internal/controllers/grpc/films"
	clientgrpc "service-2/internal/infrastructure/grpc"
	"service-2/internal/usecases/persistent"

	"google.golang.org/grpc"
)

func Run() {
	filmsServiceAddr := fmt.Sprintf("%s:%s",
		os.Getenv("KONG_HOST"),
		os.Getenv("KONG_PORT"))

	filmsClientFactory := clientgrpc.NewFilmsClientsFactory(filmsServiceAddr)

	usecase := persistent.NewFilmsUc(filmsClientFactory)
	gRPCServer := grpc.NewServer()
	filmsgrpc.Register(gRPCServer, usecase)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
	if err != nil {
		log.Fatalf("listen error: %s", err)
	}

	log.Println("Server starting...")
	if err = gRPCServer.Serve(l); err != nil {
		log.Fatalf("start serve error: %s", err)
	}
}
