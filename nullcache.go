package abreuvoir

// NullCache discards anything attempting to be cached.
type NullCache struct{}

func loadCache(host string) *NullCache {
	return &NullCache{}
}

func (cache *NullCache) cacheData(key string, data []byte) {

}
