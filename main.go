package main

import (
	"context"
	"fmt"

	"github.com/MauroMontan/grafito/grafito"
)

type Recipe struct {
	Name   string `json:"name"`
	Asset  string `json:"asset"`
	Spoils string `json:"spoils"`
}

type Response struct {
	Data struct {
		CrockpotRecipes []Recipe `json:"crockpotRecipes"`
	} `json:"data"`
}

func main() {

	url := "https://graphql.dont-starve-together-api.xyz/"

	client := grafito.NewClient(url)

	client.AddHeader("poweredBy", "Helado4Night").AddHeader("hello", "world")

	q := grafito.Query{
		Name:      "character",
		Arguments: map[string]any{},
		Fields: []string{
			"name",
			"asset",
			"spoils",
		},
	}

	data := &Response{}

	ctx := context.TODO()

	err := client.RunQuery(ctx, q, data)

	if err != nil {
		println("cannot query")
	}

	for _, v := range data.Data.CrockpotRecipes {
		fmt.Printf("v.Name: %v\n", v.Spoils)
	}

}
