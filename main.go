package main

import (
	"log"
	"net"
	"user-service/proto"
	"user-service/service"

	emailproto "github.com/Ukkenjijo/email-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	emailConn, _ := grpc.NewClient("email-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials())) // Connect to Email Service
	emailClient := emailproto.NewEmailServiceClient(emailConn)

	proto.RegisterUserServiceServer(grpcServer, &service.UserService{EmailClient: emailClient})

	log.Println("User Service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
