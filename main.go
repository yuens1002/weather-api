package main

import (
	"io/ioutil"

	graphql "github.com/graph-gophers/graphql-go"
	c "github.com/yuens1002/server-graphql-go-weather/config"
	"github.com/yuens1002/server-graphql-go-weather/server"
	util "github.com/yuens1002/server-graphql-go-weather/utils"

	_ "github.com/lib/pq"
)

func main() {
	util.ViperInt()
	dbCfg := c.DBconfig()
	cs := ConnString(dbCfg)
	db := ConnectToDB(cs)
	defer db.Close()

	// Parse Schema
	bStr, err := ioutil.ReadFile("./schema.graphql")
	util.Check(err, "ioutil.ReadFile")
	schemaString := string(bStr)
	var Schema *graphql.Schema
	Schema, err = graphql.ParseSchema(schemaString, &RootResolver{db})
	util.Check(err, "graphql.ParseSchema")

	// start mux
	server.StartServer(Schema)

}
