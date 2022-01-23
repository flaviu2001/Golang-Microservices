FROM golang:1.17.6

WORKDIR /app

# This assumes the context to be the parent of port-domain-service
COPY . .

RUN apt-get update
RUN apt install -y protobuf-compiler

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

CMD ["go", "run", "./port-domain-service/"]
