package main

import (
	"context"
	"fmt"

	"github.com/MauroMontan/grafito/grafito"
)

type Character struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {

	url := "https://rickandmortyapi.com/graphql"

	client := grafito.NewClient(url, grafito.HttpDefaultClient)

	client.AddHeader("poweredBy", "Helado4Night").AddHeader("hello", "world")

	q := grafito.Query{
		Name:      "characters",
		Arguments: map[string]any{},
		Fields: []string{
			"results{name}",
		},
	}

	var resp struct {
		Res struct {
			Results []Character `json:"results"`
		} `json:"characters"`
	}

	ctx := context.Background()

	err := client.RunQuery(ctx, q, &resp)

	if err != nil {
		println("cannot query")
	}

	for _, char := range resp.Res.Results {

		fmt.Printf("char: %v\n", char.Name)

	}

}
