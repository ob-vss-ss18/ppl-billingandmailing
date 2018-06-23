package api

import (
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"log"
	"github.com/graphql-go/graphql"
)

func Init_graphql() *gqlhandler.Handler{
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query,
		Mutation: mutation,
	})
	if err != nil {
		log.Fatalf("scheme creation failed with error: %v", err)
	}
	graphqlhandler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
	return graphqlhandler
}
