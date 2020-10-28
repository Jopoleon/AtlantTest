## Atlant Test Task 


### Requirements
- Go version > 1.14
- Docker
- protobuf (protoc compiler)
### Before start

 - `.env`  file with configuration values 
 
    ```.env
    HTTP_PORT=8080
    DB_HOST=localhost
    DB_USER=mongodb
    DB_PASSWORD=mongodb
    DB_NAME=products
    DB_PORT=27017
    ```
 - `go mod vendor` for downloading dependencies (only for local start)
 
 - generate `proto_server` `protoc -I=proto\ --go_out=plugins=grpc:server proto/*.proto` 

### How to start

#### Local
  
  - `go build -o atalant `
  - `./atlant`
  
#### Docker 
  You can use `--scale=n` flag with `docker-compose` 
  command to start `n` number of containers and then `nginx` will do the balancing.

 - `docker-compose up --build --scale atlant_test=3`