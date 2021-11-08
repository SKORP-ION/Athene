package main

import (
	"Athena/ldap"
	"Athena/ldap/api"
	"Athena/ldap/service"
	. "Athena/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	service := service.GrpcService{}

	api.RegisterLdapServer(server, service)
	Info.Println("Service is registered")

	conn, err := net.Listen("tcp", ldap.ServiceHost)

	if err != nil {
		Error.Fatalln(err)
	}

	Info.Printf("Create connection at %s\n", ldap.ServiceHost)
	err = server.Serve(conn)

	if err != nil {
		Error.Fatalln(err)
	}

}
