.PHONY: postgres adminer migrate server

server:
	go run .\server\main.go

migrate:
	migrate -source file://migrations -database postgres://postgres:12345@localhost/pharmacy?sslmode=disable up

migrate-down:
	migrate -source file://migrations -database postgres://postgres:12345@localhost/pharmacy?sslmode=disable down