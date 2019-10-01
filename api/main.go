package main

import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

var Schema = `
	schema {
		query: Query
	}
	type Person{
		id: ID!
		firstName: String!
		lastName: String
	}
	type Query{
		person(id: ID!): Person
	}
`

type personResolver struct {
	p *Person
}

func (r *personResolver) ID() graphql.ID {
	return r.p.ID
}

func (r *personResolver) FirstName() string {
	return r.p.FirstName
}

func (r *personResolver) LastName() *string {
	return &r.p.LastName
}

type Resolver struct{}

func (r *Resolver) Person(args struct{ ID graphql.ID }) *personResolver {
	if p := peopleData[args.ID]; p != nil {
		log.Print("Found in resolver!/n")
		return &personResolver{p}
	}
	return nil
}

var peopleData = make(map[graphql.ID]*Person)

var mainSchema *graphql.Schema

func init() {
	for _, p := range People {
		peopleData[p.ID] = p
	}

	mainSchema = graphql.MustParseSchema(Schema, &Resolver{})
}

func main() {
	http.Handle("/query", &relay.Handler{Schema: mainSchema})
	log.Print("Starting to listen 8964...")
	log.Fatal(http.ListenAndServe(":8964", nil))
}
