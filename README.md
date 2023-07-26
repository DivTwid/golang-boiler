# golang-boiler

This is a basic Go lang boiler plate with Gin-gonic and GORM implementation

To Start with the mentioned boiler plate simply clone the repository and run following commands
`ENV=development go run main.go`

Project Overview:

    - main.go
    - config.development.yaml
    - config.production.yaml
    - Dockerfile
    - Makefile
    - README.md
    - config/
        - config.go
    - controller/
     - DummyController.go
    - db/
        - migration/
          - migrate.go
        - seeders
          - UserSeeders.go
        - db.go
    - dto
        - UserDto.go
    - env
        - .env_test
    - middleware
        - basicAuth.go
    - model
        - users.go
    - router
        - routes.go
    - service
        - DummyService.go
    - templates
        - index.tmpl

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



