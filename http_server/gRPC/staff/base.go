package staff

import (
	"Athena/staff/api"
	"google.golang.org/grpc"
	"os"
)

var Client api.StaffClient

func init() {
	host := os.Getenv("staff_host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	Client = api.NewStaffClient(conn)
}
