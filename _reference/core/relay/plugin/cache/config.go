package cache

type Config struct {
	Enable            bool   `json:"enable"`
	TTL               int    `json:"ttl"`
	ItemMaxSize       int    `json:"item_max_size"`
	AddCacheHitHeader bool   `json:"add_cache_hit_header"`
	CacheHitHeader    string `json:"cache_hit_header"`
}
