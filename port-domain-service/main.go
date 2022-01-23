package main

import (
	"Bleenco/common"
	pb "Bleenco/rpc"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net"
)

const (
	createAliasTable = "CREATE TABLE IF NOT EXISTS aliases (" +
		"id SERIAL PRIMARY KEY," +
		"port_id BIGINT REFERENCES ports(id)," +
		"unlocs VARCHAR," +
		"alias VARCHAR)"
	createRegionTable = "CREATE TABLE IF NOT EXISTS regions (" +
		"id SERIAL PRIMARY KEY," +
		"port_id BIGINT REFERENCES ports(id)," +
		"unlocs VARCHAR," +
		"region VARCHAR)"
	createPortTable = "CREATE TABLE IF NOT EXISTS ports (" +
		"id BIGINT PRIMARY KEY," +
		"unlocs VARCHAR UNIQUE," +
		"name VARCHAR," +
		"city VARCHAR," +
		"country VARCHAR," +
		"coord1 REAL," +
		"coord2 REAL," +
		"province VARCHAR," +
		"timezone VARCHAR," +
		"code VARCHAR)"
	removeAliases       = "DELETE FROM aliases WHERE unlocs = $1"
	removeRegions       = "DELETE FROM regions WHERE unlocs = $1"
	selectHighestId     = "SELECT GREATEST(0, max(id)) from (select id from ports order by id desc limit 1) t"
	selectPortId        = "SELECT id FROM ports WHERE unlocs = $1"
	upsertPortStatement = "INSERT INTO ports " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) " +
		"ON CONFLICT (unlocs) DO UPDATE " +
		"SET name = $3, city = $4, country = $5, coord1 = $6, coord2 = $7, province = $8, timezone = $9, code = $10"
	insertAlias         = "INSERT INTO aliases(port_id, unlocs, alias) values ($1, $2, $3)"
	insertRegion        = "INSERT INTO regions(port_id, unlocs, region) values ($1, $2, $3)"
	paginatedSelectPort = "SELECT * FROM ports WHERE id BETWEEN $1 AND $2"
	selectAliases       = "SELECT alias FROM aliases WHERE unlocs = $1"
	selectRegions       = "SELECT region FROM regions WHERE unlocs = $1"
)

type server struct {
	pb.UnimplementedCommunicatorServer
}

func (s *server) Upsert(stream pb.Communicator_UpsertServer) error {
	for {
		rpcPort, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}
		if err != nil {
			return err
		}
		upsertPort(common.RpcPortToJsonPort(rpcPort))
	}
}

func (s *server) Select(rpcPage *pb.RpcPage, stream pb.Communicator_SelectServer) error {
	ports := getPorts(int(rpcPage.Page))
	for _, port := range ports {
		if err := stream.Send(common.JsonPortToRpcPort(port)); err != nil {
			return err
		}
	}
	return nil
}

func getConnection() *sql.DB {
	host := common.FromEnvVar(common.EnvDbHost, common.DbHost)
	port := common.FromEnvVar(common.EnvDbPort, common.DbPort)
	user := common.FromEnvVar(common.EnvDbUser, common.DbUser)
	password := common.FromEnvVar(common.EnvDbPass, common.DbPassword)
	dbname := common.FromEnvVar(common.EnvDbName, common.DbName)
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return db
}

func initDatabase() {
	conn := getConnection()
	defer func(conn *sql.DB) {
		err := conn.Close()
		common.CheckError(err)
	}(conn)
	_, err := conn.Exec(createPortTable)
	common.CheckError(err)
	_, err = conn.Exec(createAliasTable)
	common.CheckError(err)
	_, err = conn.Exec(createRegionTable)
	common.CheckError(err)
}

func upsertPort(port common.Port) {
	conn := getConnection()
	defer func(conn *sql.DB) {
		err := conn.Close()
		common.CheckError(err)
	}(conn)
	unlocs := port.Unlocs[0]
	_, err := conn.Exec(removeAliases, unlocs)
	common.CheckError(err)
	_, err = conn.Exec(removeRegions, unlocs)
	common.CheckError(err)
	var coord1 interface{}
	var coord2 interface{}
	if len(port.Coordinates) != 0 {
		coord1 = port.Coordinates[0]
		coord2 = port.Coordinates[1]
	} else {
		coord1 = nil
		coord2 = nil
	}
	rows, err := conn.Query(selectHighestId)
	common.CheckError(err)
	var portId int64
	for rows.Next() {
		err = rows.Scan(&portId)
		common.CheckError(err)
	}
	err = rows.Close()
	common.CheckError(err)
	portId += 1
	_, err = conn.Exec(upsertPortStatement, portId, unlocs, port.Name, port.City, port.Country, coord1, coord2, port.Province, port.Timezone, port.Code)
	common.CheckError(err)
	rows, err = conn.Query(selectPortId, unlocs)
	common.CheckError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		common.CheckError(err)
	}(rows)
	for rows.Next() {
		err = rows.Scan(&portId)
		common.CheckError(err)
	}
	for _, alias := range port.Alias {
		_, err = conn.Exec(insertAlias, portId, unlocs, alias)
		common.CheckError(err)
	}
	for _, region := range port.Regions {
		_, err = conn.Exec(insertRegion, portId, unlocs, region)
		common.CheckError(err)
	}

}

func getAliases(conn *sql.DB, unlocs string) []string {
	rows, err := conn.Query(selectAliases, unlocs)
	common.CheckError(err)
	aliases := make([]string, 0)
	for rows.Next() {
		var alias string
		err := rows.Scan(&alias)
		common.CheckError(err)
		aliases = append(aliases, alias)
	}
	return aliases
}

func getRegions(conn *sql.DB, unlocs string) []string {
	rows, err := conn.Query(selectRegions, unlocs)
	common.CheckError(err)
	regions := make([]string, 0)
	for rows.Next() {
		var region string
		err := rows.Scan(&region)
		common.CheckError(err)
		regions = append(regions, region)
	}
	return regions
}

func getPorts(page int) []common.Port {
	conn := getConnection()
	defer func(conn *sql.DB) {
		err := conn.Close()
		common.CheckError(err)
	}(conn)
	lowerBound := page * common.PageSize
	upperBound := (page+1)*common.PageSize - 1
	rows, err := conn.Query(paginatedSelectPort, lowerBound, upperBound)
	common.CheckError(err)
	ports := make([]common.Port, 0)
	for rows.Next() {
		var id int64
		var unlocs string
		var name string
		var city string
		var country string
		var coord1 *float32
		var coord2 *float32
		var province string
		var timezone string
		var code *string
		err := rows.Scan(&id, &unlocs, &name, &city, &country, &coord1, &coord2, &province, &timezone, &code)
		common.CheckError(err)
		var coordinates []float32
		if coord1 == nil {
			coordinates = []float32{}
		} else {
			coordinates = []float32{*coord1, *coord2}
		}
		var dereferencedCode = ""
		if code != nil {
			dereferencedCode = *code
		}

		ports = append(ports, common.Port{
			Name:        name,
			City:        city,
			Country:     country,
			Alias:       getAliases(conn, unlocs),
			Regions:     getRegions(conn, unlocs),
			Coordinates: coordinates,
			Province:    province,
			Timezone:    timezone,
			Unlocs:      []string{unlocs},
			Code:        dereferencedCode,
		})
	}
	return ports
}

func main() {
	initDatabase()
	var port = common.FromEnvVar(common.GrpcServerPort, common.DefaultPort)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommunicatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
