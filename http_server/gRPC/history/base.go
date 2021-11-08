package history

import (
	"Athena/history/api"
	"google.golang.org/grpc"
	"os"
)

var Client api.ActionsAndHistoryClient

func init() {
	host := os.Getenv("history_host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	Client = api.NewActionsAndHistoryClient(conn)
}
