package cacheoverlay

// Close shuts down the cache
func (co *CacheOverlay) Close() {
	co.cache.Close()
}
