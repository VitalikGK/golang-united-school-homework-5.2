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

var ListCache []string

func NewCache() Cache {

	c := new(Cache)
	return *c
}

func (c Cache) Get(key string) (string, bool) {
	var k string
	var ok bool
	fmt.Println("Get Time = ", time.Until(c.deadline).Milliseconds())
	if c.key == key && time.Until(c.deadline).Milliseconds() <= 0 {

		k = c.value
		ok = true

	} else {
		if k != key {
			k = ""
		} else {
			k = key
		}
		ok = false
	}
	return k, ok
}

func (c *Cache) Put(key, value string) {
	c.key = key
	c.value = value
	// c.stings = time.Now().Format("Jan _2 15:04:05.000000")
	// c.deadline = time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC)
	// c.dead = false

}

func (c Cache) Keys() []string {
	// var listCache []string
	fmt.Println("Get Time = ", time.Until(c.deadline).Milliseconds())
	// for i := 0; i < len(ListCache); i++{
	if time.Until(c.deadline).Milliseconds() <= 0 {
		ListCache = append(ListCache, c.key)
	}
	return ListCache
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	m := c.deadline
	if m.Sub(deadline) > 0 {
		ListCache = append(ListCache, c.key)
	}
	fmt.Println("Keys = ", ListCache)

}
