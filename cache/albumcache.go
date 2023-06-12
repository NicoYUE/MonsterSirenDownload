package cache

import (
	"encoding/gob"
	"os"
	"sync"
)

const DefaultPath = "./.cache.gob"

type AlbumCache struct {
	mutex    sync.RWMutex
	CacheMap map[string]map[string]bool
}

func NewAlbumCache() *AlbumCache {
	cache := &AlbumCache{
		CacheMap: make(map[string]map[string]bool),
	}

	if _, err := os.Stat(DefaultPath); os.IsNotExist(err) {
		file, _ := os.Create(DefaultPath)
		defer file.Close()
		gob.Register(map[string]bool{})
		gob.NewEncoder(file).Encode(cache.CacheMap)
	}

	rfile, _ := os.OpenFile(DefaultPath, os.O_RDONLY, 0666)
	defer rfile.Close()
	gob.Register(map[string]bool{})
	decoder := gob.NewDecoder(rfile)
	decoder.Decode(&cache.CacheMap)

	return cache
}

func (cache *AlbumCache) Cache(album string, title string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if cache.CacheMap[album] == nil {
		cache.CacheMap[album] = make(map[string]bool)
	}

	cache.CacheMap[album][title] = true

	wfile, _ := os.OpenFile(DefaultPath, os.O_WRONLY, 0666)
	defer wfile.Close()
	gob.Register(map[string]bool{})
	gob.NewEncoder(wfile).Encode(cache.CacheMap)
}

func (cache *AlbumCache) AlbumExists(album string) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	return cache.CacheMap[album] != nil
}

func (cache *AlbumCache) SongExists(album string, title string) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if cache.CacheMap[album] != nil {
		return cache.CacheMap[album][title]
	}
	return false
}
