package service

import (
	"context"
	"log"
	"user-service/proto"
	emailproto "github.com/Ukkenjijo/email-service/proto"
	
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	EmailClient emailproto.EmailServiceClient
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
}

func (s *UserService) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	// Store user in DB (mocking)
	newUser := User{Username: req.Username, Email: req.Email, Password: req.Password}
	log.Printf("User registered: %+v\n", newUser)

	// Send email
	emailReq := &emailproto.EmailRequest{
		To:      req.Email,
		Subject: "Welcome to Our Platform",
		Body:    "Hello " + req.Username + ", your account has been created!",
	}

	_, err := s.EmailClient.SendEmail(ctx, emailReq)
	if err != nil {
		log.Println("Failed to send email:", err)
	}

	return &proto.RegisterResponse{Message: "User registered successfully"}, nil
}
