package main

import (
	"Bleenco/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host             = "localhost"
	port             = 5432
	user             = "postgres"
	password         = "bunica"
	dbname           = "bleenco"
	createAliasTable = "CREATE TABLE IF NOT EXISTS alias (" +
		"id SERIAL PRIMARY KEY," +
		"unlocs VARCHAR REFERENCES port(unlocs)," +
		"alias VARCHAR)"
	createRegionTable = "CREATE TABLE IF NOT EXISTS region (" +
		"id SERIAL PRIMARY KEY," +
		"unlocs VARCHAR REFERENCES port(unlocs)," +
		"region VARCHAR)"
	createPortTable = "CREATE TABLE IF NOT EXISTS port (" +
		"unlocs VARCHAR PRIMARY KEY," +
		"name VARCHAR," +
		"city VARCHAR," +
		"country VARCHAR," +
		"coord1 REAL," +
		"coord2 REAL," +
		"province VARCHAR," +
		"timezone VARCHAR)"
	removeAliases       = "DELETE FROM alias WHERE unlocs = $1"
	removeRegions       = "DELETE FROM region WHERE unlocs = $1"
	upsertPortStatement = "INSERT INTO port " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) " +
		"ON CONFLICT (unlocs) DO UPDATE " +
		"SET name = $2, city = $3, country = $4, coord1 = $5, coord2 = $6, province = $7, timezone = $8"
	insertAlias  = "INSERT INTO alias(unlocs, alias) values ($1, $2)"
	insertRegion = "INSERT INTO region(unlocs, region) values ($1, $2)"
)

func getConnection() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
		utils.CheckError(err)
	}(conn)
	_, err := conn.Exec(createPortTable)
	utils.CheckError(err)
	_, err = conn.Exec(createAliasTable)
	utils.CheckError(err)
	_, err = conn.Exec(createRegionTable)
	utils.CheckError(err)
}

func upsertPort(port utils.Port) {
	conn := getConnection()
	defer func(conn *sql.DB) {
		err := conn.Close()
		utils.CheckError(err)
	}(conn)
	unlocs := port.Unlocs[0]
	_, err := conn.Exec(removeAliases, unlocs)
	utils.CheckError(err)
	_, err = conn.Exec(removeRegions, unlocs)
	utils.CheckError(err)
	var coord1 interface{}
	var coord2 interface{}
	if len(port.Coordinates) != 0 {
		coord1 = port.Coordinates[0]
		coord2 = port.Coordinates[1]
	} else {
		coord1 = nil
		coord2 = nil
	}
	_, err = conn.Exec(upsertPortStatement, unlocs, port.Name, port.City, port.Country, coord1, coord2, port.Province, port.Timezone)
	for _, alias := range port.Alias {
		_, err = conn.Exec(insertAlias, unlocs, alias)
		utils.CheckError(err)
	}
	for _, region := range port.Regions {
		_, err = conn.Exec(insertRegion, unlocs, region)
		utils.CheckError(err)
	}

}

func handleUpsert(w http.ResponseWriter, r *http.Request) {
	var port utils.Port
	_ = json.NewDecoder(r.Body).Decode(&port)
	upsertPort(port)
	var response = utils.JsonStatusResponse{Status: "success"}
	err := json.NewEncoder(w).Encode(response)
	utils.CheckError(err)
}

func main() {
	initDatabase()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/upsert", handleUpsert).Methods("POST")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
