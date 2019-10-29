package cacheoverlay

import (
	"errors"
	"github.com/dgraph-io/ristretto"
)

type (
	CacheOverlay struct {
		cache     *ristretto.Cache
		retriever Retrieve
	}

	Retrieve func(key string) (interface{}, error)
)

// NewCacheOverlay returns an initialized cache overlay object
func NewCacheOverlay(config *ristretto.Config, retriever Retrieve) (*CacheOverlay, error) {
	if retriever == nil {
		return nil, errors.New("no retriever provided")
	}
	c, err := ristretto.NewCache(config)
	if err != nil {
		return nil, err
	}
	return &CacheOverlay{
		cache:     c,
		retriever: retriever,
	}, nil
}
