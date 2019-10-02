/*
Getting Started With GraphQL Using Golang
https://www.thepolyglotdeveloper.com/2018/05/getting-started-graphql-golang/
*/
package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{})

	// Create a http handler at :8964/graphql
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8964", nil)
}

/*
The /graphql endpoint will accept requests that look like this:
	http://localhost:8964/graphql?query={aaa}

	The query property {aaa} represents the GraphQL query that we’re sending to the server.
	This dummy GraphQL api returns empty result
	{"data":{}}
*/
