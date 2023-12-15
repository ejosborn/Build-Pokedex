package pokeapi

import (
	"net/http"
	"time"

	"github.com/ejosborn/Build-Pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(inCache, wait time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(inCache),
		httpClient: http.Client{
			Timeout: wait,
		},
	}
}
