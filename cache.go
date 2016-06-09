package ecosystem

import (
	"time"

	"github.com/diegobernardes/ttlcache"
)

//EmailPWCache is the cache for storing email/temp pw combinations for passwordless authorisation
var EmailPWCache = initCache(300)

func initCache(exp time.Duration) *ttlcache.Cache {
	newCache := ttlcache.NewCache()
	newCache.SetTTL(time.Duration(exp * time.Second))
	return newCache
}
