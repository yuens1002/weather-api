package server

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/rs/cors"
	"github.com/spf13/viper"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/yuens1002/server-graphql-go-weather/handler"
	util "github.com/yuens1002/server-graphql-go-weather/utils"
)

// StartServer starts the mux server with customizable schema and port
func StartServer(schema *graphql.Schema) {
	ep := "/graphql"
	defaultPort := viper.GetString("local-port")
	str := viper.GetString("app-instance")
	p := viper.GetString("PORT")
	if p == "" {
		p = defaultPort
		str = fmt.Sprintf("http://localhost:%s", p)
	}
	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/weather", &handler.CityState{})
	mux.Handle(ep, &relay.Handler{Schema: schema})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", p),
		Handler: cors.Default().Handler(mux),
	}

	color.Blue("GraphQL started at %s%s \n", str, ep)
	err := s.ListenAndServe()
	util.Check(err, "ListenAndServe")

}
