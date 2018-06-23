package api

import "github.com/graphql-go/graphql"


var contractType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Leasing_contract",
		Fields: graphql.InputObjectConfigFieldMap{
			"leasing_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"kunden_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"products": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.Int),
			},
			"datum": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"rabatt": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"service_flat": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"testwert": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"versicherung": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
	},
)
