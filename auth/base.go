package auth

import (
	. "Athena/log"
	"os"
)

var (
	dbUser      = ""
	dbPassword  = ""
	dbName      = ""
	dbHost      = ""
	dbPort      = ""
	ServiceHost = ""
)

func init() {
	dbUser = os.Getenv("dbUser")
	dbPassword = os.Getenv("dbPassword")
	dbName = os.Getenv("dbName")
	dbHost = os.Getenv("dbHost")
	dbPort = os.Getenv("dbPort")
	ServiceHost = os.Getenv("serviceHost")

	err := connectDB(dbUser, dbPassword, dbHost, dbPort, dbName)

	if err != nil {
		Error.Fatalln("Can't connect to database. Error: ", err)
	}
}
