package cache

import (
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
	k := c.key + ", " + c.value + ", " + c.stings
	return k, c.dead
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
