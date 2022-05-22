package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	Kv       map[string]string
	Ev       map[string]string
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
		Ev:       make(map[string]string),
		Deadline: make(map[string]time.Time),
		Dead:     make(map[string]bool),
	}

	return c
}

func (c Cache) Get(key string) (string, bool) {
	var k string
	var ok bool
	fmt.Println("Get Time = ", time.Until(c.Deadline[key]).Milliseconds())
	_, exists := c.Kv[key]
	if !exists {
		// if k != key {
		k = ""
		// } else {
		// 	k = key
		// }
		ok = false
	} else {

		if time.Until(c.Deadline[key]).Milliseconds() <= 0 || c.Deadline[key].IsZero() {

			k = c.Kv[key]
			ok = true

		} else {
			k = c.Kv[key]
			ok = false

		}

	}
	return k, ok
}

func (c Cache) Put(key, value string) {
	fmt.Println("c.Kv[key],  value", c.Kv[key], value)
	_, exists := c.Kv[key]
	if !exists {
		//if KV[key] == key {
		c.Kv[key] = value
		c.Ev[key] = value
		c.Deadline[key] = time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC)
		c.Dead[key] = false
		fmt.Println("Добавили Cache ", c.Kv[key], c.Deadline[key], c.Dead[key])
	} else {
		c.Kv[key] = value
		c.Ev[key] = value
		c.Deadline[key] = time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC)
		c.Dead[key] = false
		fmt.Println("Изменили Cache ", c.Kv[key], c.Deadline[key], c.Dead[key])
	}

}

func (c Cache) Keys() []string {
	var listCache []string
	for i, key := range c.Kv {
		fmt.Println("Get Time = ", time.Until(c.Deadline[i]).Milliseconds())
		fmt.Println("Get Time key = ", key, i)
		//if time.Until(key).Milliseconds() <= 0 {
		//		if key != time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC) {
		fmt.Println("___________________________________________________________________________")
		fmt.Println("Cache ", c)

		fmt.Println("___________________________________________________________________________")
		if time.Until(c.Deadline[i]).Milliseconds() >= 0 || c.Deadline[i].IsZero() {
			listCache = append(listCache, i)
			//			}
		}
	}
	fmt.Println("listCache = ", listCache)
	return listCache
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	fmt.Println("c.Kv[key],  value", c.Kv[key], value)
	_, exists := c.Kv[key]
	if !exists {
		//if KV[key] == key {
		c.Kv[key] = value
		c.Ev[key] = value
		c.Deadline[key] = deadline
		c.Dead[key] = false
		fmt.Println("Добавили Cache ", c.Kv[key], c.Deadline[key], c.Dead[key])
	} else {
		c.Kv[key] = value
		c.Ev[key] = value
		c.Deadline[key] = deadline
		c.Dead[key] = false
		fmt.Println("Изменили Cache ", c.Kv[key], c.Deadline[key], c.Dead[key])
	}
}
