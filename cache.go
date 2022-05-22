package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	Kv       map[string]string
	Deadline map[string]time.Time
	Dead     map[string]bool
}

var KV map[string]string = make(map[string]string)
var DD map[string]time.Time = make(map[string]time.Time)
var DaD map[string]bool = make(map[string]bool)

//type MapCache map[string]Cache

func NewCache() Cache {
	c := Cache{
		Kv:       make(map[string]string),
		Deadline: make(map[string]time.Time),
		Dead:     make(map[string]bool),
	}

	return c
}

func (c Cache) Get(key string) (string, bool) {
	var k string
	var ok bool
	fmt.Println("Get Time = ", time.Until(DD[key]).Milliseconds())
	_, exists := DD[key]
	if !exists {
		if k != key {
			k = ""
		} else {
			k = key
		}
		ok = false
	} else {

		if time.Until(DD[key]).Milliseconds() <= 0 {

			k = KV[key]
			ok = true

		}

	}
	return k, ok
}

func (c Cache) Put(key, value string) {
	fmt.Println("KV[key],  value", KV[key], value)
	_, exists := DD[key]
	if !exists {
		//if KV[key] == key {
		KV[key] = value
		DD[key] = time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC)
		DaD[key] = false
		fmt.Println("Добавили Cache ", KV[key], DD[key], DaD[key])
	} else {
		KV[key] = value
		DD[key] = time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC)
		DaD[key] = false
		fmt.Println("Изменили Cache ", KV[key], DD[key], DaD[key])
	}

}

func (c Cache) Keys() []string {
	var listCache []string
	for i, key := range DD {
		fmt.Println("Get Time = ", time.Until(key).Milliseconds())
		if time.Until(key).Milliseconds() <= 0 {
			listCache = append(listCache, i)
		}
	}
	return listCache
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	var listCache []string
	for i, key := range DD {
		fmt.Println("PutTill Get Time = ", time.Until(key).Milliseconds())
		if key.Sub(deadline) <= 0 {
			listCache = append(listCache, i)
		}
	}
	fmt.Println("listcache = ", listCache)
}
