package main

import (
	"context"
	"time"

	"github.com/MauroMontan/grafito/grafito"
)

type Character struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type RickAndMortyService struct {
	client *grafito.Client
}

func NewRickAndMortyService() *RickAndMortyService {

	c := &grafito.Client{
		Url:  "https://rickandmortyapi.com/graphql",
		Http: grafito.HttpDefaultClient,
	}

	return &RickAndMortyService{client: c}

}

func (s *RickAndMortyService) GetCharacters(ctx context.Context, page int) ([]Character, error) {
	q := grafito.Query{
		Name: "characters",
		Arguments: map[string]any{
			"page": page,
		},
		Fields: []string{
			"results{id name status}",
		},
	}

	var resp struct {
		Characters struct {
			Results []Character `json:"results"`
		} `json:"characters"`
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err := s.client.RunQuery(ctx, q, &resp); err != nil {
		cancel()
		return nil, err
	}

	return resp.Characters.Results, nil
}
