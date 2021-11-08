package auth

import (
	"Athena/auth/api"
	"google.golang.org/grpc"
	"os"
)

var Client api.AuthorizationClient

func init() {
	host := os.Getenv("auth_host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	Client = api.NewAuthorizationClient(conn)
}
