package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	key      string
	value    string
	stings   string
	deadline time.Time
	dead     bool
}

type MapCache map[string]Cache

func NewCache() Cache {

	c := new(Cache)
	return *c
}

func (c Cache) Get(key string) (string, bool) {
	k := c.value
	ok := c.dead
	if k != key {
		k = ""
		ok = false
	}
	return k, ok
}

func (c *Cache) Put(key, value string) {
	c.key = key
	c.value = value
	c.stings = time.Now().Format("Jan _2 15:04:05.000000")
	c.deadline = time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC)
	c.dead = false

}

func (c Cache) Keys() []string {
	var listCache []string
	if c.deadline != time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC) {
		listCache = append(listCache, c.key+", "+c.value+", "+c.stings+", "+c.deadline.Format("15:04:05"))
	}
	return listCache
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	var listCache []string
	if c.deadline.Sub(deadline) >= 0 {
		listCache = append(listCache, c.key+", "+c.value+", "+c.stings+", "+c.deadline.Format("15:04:05"))
	}
	fmt.Println("Keys = ", listCache)

}
