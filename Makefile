
Server:
	go run server/server.go 
Client:
	go run client/client.go 
Migration:
	cd migration && goose postgres "postgres://mydb1:123456@localhost:5432/mydb1?sslmode=disable" up
Immigration:
	cd migration && goose postgres "postgres://mydb1:123456@localhost:5432/mydb1?sslmode=disable" down
	
