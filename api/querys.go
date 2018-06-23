package api

import (
	"github.com/graphql-go/graphql"
	"log"
)

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getBilling" : &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams)(interface{}, error){
				log.Println("Query recieved")
				return "Operation not supported!", nil
			},
		},
	},
})