package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"banks-api/config"
	appbanks "banks-api/internal/app/banks"
	"banks-api/internal/pkg/database"
	"banks-api/internal/pkg/logger"
	"banks-api/internal/pkg/server"
	"banks-api/internal/pkg/services/banks"
	pbbanks "banks-api/pkg/banks"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort = "8080"
	grpcPort = "8081"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("error load configs:", err)
	}

	logger.NewLogger(logger.DebugLogLevel)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", grpcPort))
	if err != nil {
		log.Fatalf("error when listen grpc: %s", err.Error())
	}
	defer listener.Close()

	pool, err := pgxpool.Connect(ctx, cfg.GetCheckoutPgConnectionString())
	if err != nil {
		logger.Fatalf("connect to db: %s", err.Error())
	}
	defer pool.Close()

	banksRepo := database.NewBanksRepository(pool)

	banksService := banks.NewBanksService(banksRepo)

	grpcServer := grpc.NewServer()
	banksHandler := appbanks.NewBanksHandler(banksService)
	pbbanks.RegisterBanksServer(grpcServer, banksHandler)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed when serve: %s", err.Error())
		}
	}()

	workerConn, err := grpc.NewClient(
		listener.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %s", err.Error())
	}
	defer workerConn.Close()

	mux := runtime.NewServeMux()
	if err = pbbanks.RegisterBanksHandler(ctx, mux, workerConn); err != nil {
		log.Fatalf("failed to register handler: %s", err.Error())
	}

	s := server.NewHTTPServer(mux, httpPort)
	logger.Infof("launched server at", httpPort)
	log.Fatal(s.Launch())
}
