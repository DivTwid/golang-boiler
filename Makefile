runDB:
	docker run --name postgres-sql -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres

installCompileDaemon:
	go get github.com/githubnemo/CompileDaemon

installPostgresDriver:
	go get gorm.io/driver/postgres

installGin:
	go get github.com/gin-gonic/gin

run:
	CompileDaemon -command="./go-boiler"

test:
	go test -v -cover ./...

DockerBuild:
	docker build -t go-boiler .

DockerRun:
	docker run -d -p 7000:4000 go-boiler

.PHONY: runDB run installPostgresDriver installCompileDaemon installGin test DockerBuild DockerRun