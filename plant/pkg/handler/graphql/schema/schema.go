package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/garhus2020/ESIhw2/plant/pkg/domain"
)

var PlantList []domain.Plant

var plantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Plant",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"ident": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"plant": &graphql.Field{
			Type: plantType,
			Description: "Get all plants",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return PlantList, nil
			},
		},
	},
})

var TodoSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
})