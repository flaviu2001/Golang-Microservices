package common

const (
	GrpcServerAddr string = "GRPC_SERVER_ADDR"
	GrpcServerPort string = "GRPC_SERVER_PORT"
	DefaultAddress string = "localhost"
	DefaultPort    string = "50051"

	EnvDbHost  string = "DBHOST"
	DbHost     string = "localhost"
	EnvDbPort  string = "DBPORT"
	DbPort     string = "5432"
	EnvDbUser  string = "DBUSER"
	DbUser     string = "postgres"
	EnvDbPass  string = "DBPASS"
	DbPassword string = "bunica"
	EnvDbName  string = "DBNAME"
	DbName     string = "bleenco"

	PortsJsonFilename     string = "ports.json"
	PortsJsonFilenameTest string = "ports_test.json"
	ChannelSize           int    = 10000
	PageSize              int    = 100
)
