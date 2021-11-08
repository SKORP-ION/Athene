package main

import (
	"Athena/auth"
	"Athena/auth/api"
	"Athena/auth/service"
	. "Athena/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	service := service.GrpcService{}

	api.RegisterAuthorizationServer(server, service)
	Info.Println("Service is registered")

	conn, err := net.Listen("tcp", auth.ServiceHost)

	if err != nil {
		Error.Fatalln(err)
	}

	Info.Printf("Create connection at %s\n", auth.ServiceHost)

	err = server.Serve(conn)

	if err != nil {
		Error.Fatalln(err)
	}

}
