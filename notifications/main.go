package main

import (
	"context"
	"fmt"
	"net"
	"service-tutorial/shared/proto"

	"google.golang.org/grpc"
)

type server struct{
	proto.UnimplementedNotificationServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err) // The port may be in use
	}
	srv := grpc.NewServer()
	proto.RegisterNotificationServiceServer(srv, &server{})
	fmt.Println("Prepare to serve")
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) SendRegistrationEmail(ctx context.Context, in *proto.RegistrationEmailRequest) (*proto.RegistrationEmailResponse, error) {
	fmt.Println("sending registration email for ", in.Email)
	return &proto.RegistrationEmailResponse{
		Success: true,
	}, nil
}
