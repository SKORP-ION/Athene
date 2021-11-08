package main

import (
	. "Athena/log"
	"Athena/staff"
	"Athena/staff/api"
	"Athena/staff/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	service := service.GrpcService{}

	api.RegisterStaffServer(server, service)
	Info.Println("Service is registered")

	conn, err := net.Listen("tcp", staff.ServiceHost)

	if err != nil {
		Error.Fatalln(err)
	}
	Info.Printf("Create connection at %s\n", staff.ServiceHost)

	err = server.Serve(conn)

	if err != nil {
		Error.Fatalln(err)
	}

}
