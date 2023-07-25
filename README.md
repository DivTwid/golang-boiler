# golang-boiler

This is a basic Go lang boiler plate with Gin-gonic and GORM implementation

To Start with the mentioned boiler plate simply clone the repository and run following commands
```go run main.go```

Project Overview:

    Main.go
        - Init
            - Initialize ENV
            - Initialize Database
            - Run Migration
            - Run Seeder
        - Main
            - Calls Gin Routes
            - Serves the project on the port mentioned in the env file


To Run project from docker

Use Following commands
```docker build -t golang-boiler .``` -> expose port 4000 from project
```docker run -d -p 4000:4000 go-boiler``` -> maps local port 4000 to docker container port 4000

Use following url to run project `localhost:4000`



