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

func NewCache(key, value, stings string, deadline time.Time, dead bool) Cache {

	return Cache{
		key:      key,
		value:    value,
		stings:   stings,
		deadline: deadline,
		dead:     dead,
	}
}

func (m MapCache) Get(key string) (string, bool) {
	k, ok := m[key]
	// if !ok {
	// 	fmt.Println("Элемента нет в списке")
	// } else {
	// 	fmt.Printf("Элемент найден (%s , %s, %s)", k.key, k.value, k.stings)
	// }
	return k.key + k.value + k.stings, ok
}

func (m MapCache) Put(key, value string) {
	k, ok := m[key]
	if ok {
		k = NewCache(key, value, k.stings, time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC), false)
		// fmt.Println("Изменили мапу ", key)
	} else {
		k = NewCache(key, value, time.Now().Format("Jan _2 15:04:05.000000"), time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC), false)
		// fmt.Println("Добавили мапу ", key)
	}
	m[key] = k

}

func (m MapCache) Keys() []string {
	var listCache []string
	for key, _ := range m {
		k, ok := m[key]
		if ok {
			if k.deadline != time.Date(1, time.January, 1, 00, 0, 0, 0, time.UTC) {
				listCache = append(listCache, k.key+", "+k.value+", "+k.stings+", "+k.deadline.Format("15:04:05"))
			}
		}

	}
	// fmt.Println("Keys = ", listCache)
	return listCache
}
