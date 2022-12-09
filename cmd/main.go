package main

import (
	"log"
	"net"

	"github.com/SaidovZohid/medium_notification_service/config"
	pb "github.com/SaidovZohid/medium_notification_service/genproto/notification_service"
	"github.com/SaidovZohid/medium_notification_service/service"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load(".")

	notificationService := service.NewNotificationService(&cfg)

	listen, err := net.Listen("tcp", cfg.GrpcPort)

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, notificationService)
	reflection.Register(s)

	log.Println("gRPC server started port in: ", cfg.GrpcPort)
	if s.Serve(listen); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
}
