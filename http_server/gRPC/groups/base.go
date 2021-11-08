package groups

import (
	"Athena/groups/api"
	"google.golang.org/grpc"
	"os"
)

var Client api.GroupsClient

func init() {
	host := os.Getenv("groups_host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	Client = api.NewGroupsClient(conn)
}
