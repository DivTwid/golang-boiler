runPostgresDB:
	docker run --name postgres-sql -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -p 5432:5432 -d postgres

runMysqlDB:
	docker run --name mysql-local -e MYSQL_ROOT_PASSWORD=root -e MYSQL_ROOT_USER=root -p 3306:3306 -d mysql:latest

installCompileDaemon:
	go get github.com/githubnemo/CompileDaemon

installPostgresDriver:
	go get gorm.io/driver/postgres

installGin:
	go get github.com/gin-gonic/gin

run:
	CompileDaemon -command="./golang-boiler"

test:
	go test -v -cover ./...

build:
	docker build -t go-boiler .

buildrun:
	docker run -d -p 8000:4000 go-boiler

.PHONY: runPostgresDB runMysqlDB run installPostgresDriver installCompileDaemon installGin test build buildrun