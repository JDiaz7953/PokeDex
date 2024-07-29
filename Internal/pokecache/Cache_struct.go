package pokecache


import (
    "time"
)


//functions
func NewCache(interval time.Duration) Cache {
    c := Cache{
        cacheData: make(map[string]CacheEntry),
    }

	go c.reapLoop(interval)
	return c
}


// Structs
type Cache struct{
	cacheData map[string]CacheEntry
}

type CacheEntry struct{
	createdAt time.Time
	val []byte
}

//Methods
func (data *Cache) Add(key string, val []byte)  {
	 data.cacheData[key] = CacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (data *Cache) Get(key string) ([]byte, bool){
	_, ok := data.cacheData[key]
	if !ok {
		return data.cacheData[key].val, false
	}
	return data.cacheData[key].val, true
}
func (data *Cache) removeOld(interval time.Duration){
	timeAgo := time.Now().UTC().Add(-interval)
	for key, val := range data.cacheData{
		if val.createdAt.Before(timeAgo){
			delete(data.cacheData,key)
		}
	}
}

func (data *Cache) reapLoop(interval time.Duration){
	ticker:= time.NewTicker(interval)
	for range ticker.C{
		data.removeOld(interval)
	}

}