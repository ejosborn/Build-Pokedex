package pokeapi

import (
	"net/http"
	"time"

	"github.com/ejosborn/Build-Pokedex/internal/pokecache"
	"github.com/ejosborn/Build-Pokedex/internal/pokedex"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache           pokecache.Cache
	personalPokedex pokedex.Pokedex
	httpClient      http.Client
}

func NewClient(inCache, wait time.Duration) Client {
	return Client{
		cache:           pokecache.NewCache(inCache),
		personalPokedex: pokedex.NewPokedex(),
		httpClient: http.Client{
			Timeout: wait,
		},
	}
}
