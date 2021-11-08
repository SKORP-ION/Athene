package ldap

import (
	"Athena/ldap/api"
	"google.golang.org/grpc"
	"os"
)

var Client api.LdapClient

func init() {
	host := os.Getenv("ldap_host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	Client = api.NewLdapClient(conn)
}
