# Golang microservices assignment

### Features
- Client and Service communication with grpc
- Postgres database for persistance
- Docker and docker-compose
- Read JSON in a buffered way to avoid loading the entire file in memory
- Pagination for data retrieval to avoid overloading the network
- Package based separation of concerns
- Goland builtin linter

### Client
The client parses a json file with ports and feeds it the server for persistence, allowing them to afterwards be retrieved. It can be configured to connect to the grpc server at an arbitrary address through environment variables. `GRPC_SERVER_ADDR` for its address and `GRPC_SERVER_PORT` for its port.

### Server
The server allows upsertion of ports and paginated retrievals through remote procedure calling. Similar to the client, the grpc port must also be specified in the `GRPC_SERVER_PORT` environment variable. Furthermore, the postgres database connection information is expected through `DBHOST`, `DBPORT`, `DBUSER`, `DBPASS` and `DBNAME`. You can find their defaults in the common/constants.go file. 

### Running
The application can be run by executing `docker-compose up`, the environment variables are set up in a way that requires no further setting up to function well. You can then test the app by running `curl localhost:8080/parse` to parse the json file and `curl localhost:8080/select/3` to retrieve the third page, or any page of data (a page contains 100 ports).

It is important to execute `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative rpc/rpc.proto` whenever you modify `rpc/rpc.proto`. This method generates the go code that allows the bridging between client and server.

As you might have noticed in the dockerfiles, the following commands are necessary for grpc to work
```
apt install -y protobuf-compiler # or the equivalent of your distro
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
```

If you're using docker you do not need to run them, but you do need to if you plan to run manually.

### Time spent
- learning golang basics: 3 hours
- port-parser branch: 2 hours
  - the part of the code that reads the json file
- microservices branch: 3 hours
  - here I structured the code to communicate through REST apis between client and server
- select-ports branch: 2 hours
  - here I implemented the paginated retrieval of data
- grpc-refactor branch: 3 hours
  - here I refactored the code to change from REST to grpc as the communication channel between client and server
- dockerize branch: 4 hours
  - here I dockerized the application and wrote the docker-compose.yaml as well
- tests and documentation: 3 hours
  - here I wrote a couple of tests and provided documentation inside the code and also wrote this README.md file