package cache

// Null discards anything attempting to be cached.
type Null struct{}

func loadCache(host string) *Null {
	return &Null{}
}

func (cache *Null) cacheData(key string, data []byte) {

}
