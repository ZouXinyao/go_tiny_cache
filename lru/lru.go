package lru

import "container/list"

type Cache struct {
	MaxEntries int
	ll *list.List
	cache map[string]*list.Element
	OnEvicted func(key string, value interface{})
}

type entry struct {
	key string
	value interface{}
}

func New(Entries int, onEvicted func(string, interface{})) *Cache {
	return &Cache{
		MaxEntries: Entries,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value interface{}, ok bool) {
	if c == nil {
		return
	}
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	if c == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
	if c.OnEvicted != nil {
		// 缓存淘汰时如果有回调函数，会直接调用。
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *Cache) Add(key string, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[string]*list.Element)
		c.ll = list.New()
	}
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

