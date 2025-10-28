package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type entry struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem

	mu sync.Mutex
}

func NewCache(capacity int) Cache {
	if capacity < 1 {
		capacity = 1
	}
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.items[key]; ok {
		node.Value = entry{key: key, value: value}
		c.queue.MoveToFront(node)
		return true
	}

	node := c.queue.PushFront(entry{key: key, value: value})
	c.items[key] = node

	if len(c.items) > c.capacity {
		lru := c.queue.Back()
		k := lru.Value.(entry).key
		c.queue.Remove(lru)
		delete(c.items, k)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.queue.MoveToFront(node)
	return node.Value.(entry).value, true
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
