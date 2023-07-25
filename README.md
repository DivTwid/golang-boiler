# golang-boiler

This is a basic Go lang boiler plate with Gin-gonic and GORM implementation

To Start with the mentioned boiler plate simply clone the repository and run following commands
`go run main.go`

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
`docker build -t golang-boiler .` -> expose port 4000 from project
`docker run -d -p 4000:4000 go-boiler` -> maps local port 4000 to docker container port 4000

Use following url to run project `localhost:4000`

Post Project is up and running the below mentioned API could be used to test the application:
`curl --location --request POST '127.0.0.1:4000/addVal' \
--header 'Authorization: Basic cm9vdDpyb290' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@tatwa.com",
    "name" : "testyUsier2",
    "phone_no": "1234567543453"
}'`


One can refer Makefile mentioned in the project to test the application test cases using below command 
`make test`



