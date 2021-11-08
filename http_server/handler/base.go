package handler

import "os"

var (
	Domain  = ""
	Host    = ""
	Cert    = ""
	PrivKey = ""
)

func init() {
	Host = os.Getenv("host")
	Domain = os.Getenv("domain")
	Cert = os.Getenv("cert")
	PrivKey = os.Getenv("privkey")
}
