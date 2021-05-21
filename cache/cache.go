package cacheDb

import "github.com/patrickmn/go-cache"

var CacheDb *cache.Cache

func InitCache() {
	CacheDb = cache.New(-1, -1)
	CacheDb.Set("usedTimes", "0", -1)
}
