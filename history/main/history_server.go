package main

import (
	"Athena/history"
	"Athena/history/api"
	"Athena/history/service"
	. "Athena/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	service := service.GrpcService{}

	api.RegisterActionsAndHistoryServer(server, service)
	Info.Println("Service is registered")

	conn, err := net.Listen("tcp", history.ServiceHost)

	if err != nil {
		Error.Fatalln(err)
	}
	Info.Printf("Create connection at %s\n", history.ServiceHost)

	err = server.Serve(conn)

	if err != nil {
		Error.Fatalln(err)
	}

}
