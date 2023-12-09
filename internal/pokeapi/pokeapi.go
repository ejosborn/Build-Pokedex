package pokeapi

import (
	"net/http"
	"time"

	"github.com/ejosborn/Build-Pokedex/internal/pokeapi"
	"github.com/ejosborn/Build-Pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache           pokecache.Cache
	personalPokedex pokeapi.Pokedex
	httpClient      http.Client
}

func NewClient(inCache, wait time.Duration) Client {
	return Client{
		cache:           pokecache.NewCache(inCache),
		personalPokedex: pokeapi.NewPokedex(),
		httpClient: http.Client{
			Timeout: wait,
		},
	}
}
