package property

import (
	"Athena/property/api"
	"google.golang.org/grpc"
	"os"
)

var Client api.PropertyClient

func init() {
	host := os.Getenv("property_host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	Client = api.NewPropertyClient(conn)
}
