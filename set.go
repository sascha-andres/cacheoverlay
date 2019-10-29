package cacheoverlay

// Set tries to write to cache
func (co *CacheOverlay) Set(key string, value interface{}) bool {
	return co.cache.Set(key, value, CacheCost)
}
