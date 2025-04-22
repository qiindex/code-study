package lru_cache

import (
	"container/list"
	"sync"
)

type entryPoint struct {
	key   int
	value int
}
type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element
	mu       sync.Mutex
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}

}

func (l *LRUCache) Get(key int) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*entryPoint).value
	}
	return -1

}

func (l *LRUCache) Put(key int, value int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if elem, ok := l.cache[key]; ok {
		elem.Value.(*entryPoint).value = value
		l.list.MoveToFront(elem)
		return
	}
	if l.list.Len() >= l.capacity {
		tail := l.list.Back()
		if tail != nil {
			delete(l.cache, tail.Value.(entryPoint).key)
		}
		l.list.Remove(tail)
	}
	elem := l.list.PushFront(&entryPoint{key, value}) //真正放
	l.cache[key] = elem

}
