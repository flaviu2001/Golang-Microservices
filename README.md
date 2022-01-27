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
The client parses a json file with ports and feeds it the server for persistence, allowing them to afterwards be retrieved. It can be configured to connect to the grpc server at an arbitrary address through environment variables.

### Server
The server allows upsertion of ports and paginated retrievals through remote procedure calling.

### Installing and running
Refer to `INSTALL.md` for the steps. 

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
  - I also refactored the code to structure the code into a repository and service pattern
- tests and documentation: 3 hours
  - here I wrote a couple of tests and provided documentation inside the code and also wrote this README.md file