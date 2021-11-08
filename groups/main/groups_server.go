package main

import (
	"Athena/groups"
	"Athena/groups/api"
	"Athena/groups/service"
	. "Athena/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	service := service.GrpcService{}

	api.RegisterGroupsServer(server, service)
	Info.Println("Service is registered")

	conn, err := net.Listen("tcp", groups.ServiceHost)

	if err != nil {
		Error.Fatalln(err)
	}
	Info.Printf("Create connection at %s\n", groups.ServiceHost)

	err = server.Serve(conn)

	if err != nil {
		Error.Fatalln(err)
	}
}
