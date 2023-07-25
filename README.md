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

