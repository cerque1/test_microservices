package app

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	filmsgrpc "service-1/internal/controllers/grpc/films"
	"service-1/internal/migrate"
	repo_persistent "service-1/internal/repo/persistent"
	usecase_persistent "service-1/internal/usecases/persistent"
	"time"

	"google.golang.org/grpc"
)

func Run() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_ADDR"),
		os.Getenv("POSTGRES_DB"))

	migrate.ApplyMigrations("file://migrations", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error open db: %s", err)
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	usecase := usecase_persistent.NewFilmsUC(repo_persistent.NewFilmsRepo(db))

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
