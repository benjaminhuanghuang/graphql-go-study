package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func main() {

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{})

	// GraphQL query, can be anything
	// To the GraphQL API, we get query from HTTP request
	query := `
    { 
			hello
    }
    `
	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}

	result := graphql.Do(params)
	rJSON, _ := json.Marshal(result)
	fmt.Printf("%s \n", rJSON) // { " data " :{}}
}
