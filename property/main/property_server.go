package main

import (
	. "Athena/log"
	"Athena/property"
	"Athena/property/api"
	"Athena/property/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	service := service.GrpcService{}

	api.RegisterPropertyServer(server, service)
	Info.Println("Service is registered")

	conn, err := net.Listen("tcp", property.ServiceHost)

	if err != nil {
		Error.Fatalln(err)
	}
	Info.Printf("Create connection at %s\n", property.ServiceHost)

	err = server.Serve(conn)

	if err != nil {
		Error.Fatalln(err)
	}

}
