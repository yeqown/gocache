package main

import (
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"log"
	"time"
)

func main() {
	pureRistretto()
}

func pureRistretto() {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000,
		MaxCost:     100,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}

	//isSet := ristrettoCache.Set("my-key", "my-value", 2)
	//fmt.Println(isSet)
	//v, ok := ristrettoCache.Get("my-key")
	//log.Println(v, ok)

	ristrettoStore := store.NewRistretto(ristrettoCache, nil)
	cacheManager := cache.New(ristrettoStore)

	err = cacheManager.Set("my-key", "my-value", &store.Options{
		Cost:       2,
		Expiration: time.Duration(5 * time.Minute),
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	value, err := cacheManager.Get("my-key")
	log.Println(value, err)
}
