package main

import (
	"log"
	"net"
	"payment-service/config"
	pb "payment-service/generated/payment_service"
	"payment-service/logs"
	"payment-service/service"
	"payment-service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	logs.InitLogger()
	logger := logs.Logger

	logger.Info("Starting the application...")

	db, err := postgres.ConnectDB()
	if err != nil {
		logger.Error("postgresga ulanishda xatolik", "error", err.Error())
		panic(err)
	}
	defer db.Close()

	config := config.Load()

	listener, err := net.Listen("tcp", config.URL_PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	s := service.NewPaymentService(*postgres.NewPaymentRepo(db), logger)
	server := grpc.NewServer()
	pb.RegisterPaymentServiceServer(server, s)

	logger.Info("server is running", "PORT", config.URL_PORT)
	log.Printf("server is running on %v...", listener.Addr())
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
