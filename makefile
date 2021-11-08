all:
	set GOOS=linux
	set GOARCH=amd64
	go build -o docker-compose\property\property property\main\property_server.go
	go build -o docker-compose\auth\auth auth\main\auth_server.go
	go build -o docker-compose\groups\groups groups\main\groups_server.go
	go build -o docker-compose\history\history history\main\history_server.go
	go build -o docker-compose\ldap\ldap ldap\main\ldap_server.go
	go build -o docker-compose\staff\staff staff\main\staff_server.go
	go build -o docker-compose\http_server\http_server http_server\http_server.go

http_server:
	set GOOS=linux
	set GOARCH=amd64
	go build -o http_server\http_server http_server\http_server.go
	xcopy http_server\static docker-compose\http_server\static
	xcopy http_server\templates docker-compose\http_server\templates
	copy http_server\favicon.ico docker-compose\http_server

property:
	set GOOS=linux
	set GOARCH=amd64
	go build -o property\main\property property\main\property_server.go

auth:
	set GOOS=linux
	set GOARCH=amd64
	go build -o auth\main\auth auth\main\auth_server.go

groups:
	set GOOS=linux
	set GOARCH=amd64
	go build -o groups\main\groups groups\main\groups_server.go

history:
	set GOOS=linux
	set GOARCH=amd64
	go build -o history\main\history history\main\history_server.go

ldap:
	set GOOS=linux
	set GOARCH=amd64
	go build -o ldap\main\ldap ldap\main\ldap_server.go

staff:
	set GOOS=linux
	set GOARCH=amd64
	go build -o staff\main\staff staff\main\staff_server.go