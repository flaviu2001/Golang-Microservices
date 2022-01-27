## Installing and running

It is important to execute `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative rpc/rpc.proto` whenever you modify `rpc/rpc.proto`. This method generates the go code that allows the bridging between client and server.

### Running with docker
This project has been tested with `Docker Compose version 2.2.3` and `Docker version 20.10.12, build e91ed5707e`.

The application can be run by executing `docker-compose up`, the environment variables are set up in a way that no further setting up is required to function well. You can then test the app by running `curl localhost:8080/parse` to parse the json file and `curl localhost:8080/select/3` to retrieve the third page, or any page of data (a page contains by default 100 ports).

### Running manually
This project has been tested with golang 1.17.6

Run the following commands to make grpc work
```
apt install -y protobuf-compiler # or the equivalent of your distro
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
```

Running the server: `go run ./port-domain-service`

Running the client: `go run ./client-api`

### Environment variables

* `GRPC_SERVER_ADDR`: address of the server
* `GRPC_SERVER_PORT`: port of the server
* `DBHOST`, `DBPORT`, `DBUSER`, `DBPASS`, `DBNAME`: address, port, username, password and database name of the used postgresql server
