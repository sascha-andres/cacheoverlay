package cacheoverlay

import "errors"

const CacheCost = 100

// Get returns an item and tries to receive when not found
func (co *CacheOverlay) Get(key string) (interface{}, error) {
	var (
		value interface{}
		found bool
		)
	if value, found = co.cache.Get(key); !found {
		newItem, err := co.retriever(key)
		if err != nil {
			return nil, err
		}
		if newItem == nil {
			return nil, nil
		}
		if !co.cache.Set(key, newItem, CacheCost) {
			return newItem, errors.New("item not added to cache")
		} else {
			return newItem, nil
		}
	}
	return value, nil
}