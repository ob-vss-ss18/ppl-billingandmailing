package api

import (
	"github.com/graphql-go/graphql"
	"log"
)

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createBilling": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"contract": &graphql.ArgumentConfig{
					Description: "contract",
					Type: contractType,
				},
			},
			Resolve: func(p graphql.ResolveParams)(interface{}, error){
				log.Println("Mutation recieved")
				// TODO: CALL CREATE BILLING FUNCTION
				return true, nil
			},
		},
	},
})
