# Golang microservices

### What's this?
As part of an assignment given to me I had to learn golang and solve a task with microservices to familiarise myself with the language and framework.

### Features
- Client and Service communication with grpc
- Postgres database for persistance
- Docker and docker-compose
- Read JSON in a buffered way to avoid loading the entire file in memory
- Pagination for data retrieval to avoid overloading the network
- Package based separation of concerns
- Goland builtin linter

### Client
The client parses a json file with ports and feeds it the server for persistence, allowing them to afterwards be retrieved. It can be configured to connect to the grpc server at an arbitrary address through environment variables.

### Server
The server allows upsertion of ports and paginated retrievals through remote procedure calling.

### Installing and running
Refer to `INSTALL.md` for the steps. 
