package main

/*
	In this demo, GraphQL handle query using schema
*/
import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			// Define resolver for filed "hello"
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "result for query 'hello'", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)

	// Query
	query := `
    { 
			hello
    }
    `
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // { "data" :{ "hello" : "result for query 'hello'"}}
}
