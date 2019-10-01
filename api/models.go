package main

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type Person struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}

var People = []*Person{
	{
		ID:        "1000",
		FirstName: "Pedro",
		LastName:  "Marquez",
	},

	{
		ID:        "1001",
		FirstName: "John",
		LastName:  "Doe",
	},
}
