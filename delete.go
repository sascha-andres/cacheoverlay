package cacheoverlay

// Delete deletes the key-value item from the cache if it exists
func (co *CacheOverlay) Delete(key string) {
	co.cache.Del(key)
}