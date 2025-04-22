package lru_cache_manually_add_lock

import (
	"container/list"
	"sync"
)

type entryPoint struct {
	key   string
	value interface{}
}
type LruCache struct {
	capacity int
	list     *list.List
	cache    map[string]*list.Element
	mu       sync.Mutex
}

func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// put
func (l *LruCache) Put(key string, value interface{}) {
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
			delete(l.cache, tail.Value.(*entryPoint).key)
		}
		l.list.Remove(tail)
	}
	elem := l.list.PushFront(&entryPoint{key, value}) //真正放
	l.cache[key] = elem

}

func (l *LruCache) Get(key string) interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*entryPoint).value
	}
	return -1
}

// 非并发
/*
func main() {
	lru := go_algo.NewLruCache(2)

	lru.Put("1", 1)

	lru.Put("2", 2)

	fmt.Println(lru.Get("1")) // 应该输出 1

	lru.Put("3", 3) // 移除键 2

	fmt.Println(lru.Get("2")) // 应该输出 -1
}

*/

// 并发

/*

func main() {
	lru := go_algo.NewLruCache(2)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		lru.Put("1", 1)
	}()

	go func() {
		defer wg.Done()
		lru.Put("2", 2)
	}()

	go func() {
		defer wg.Done()
		fmt.Println(lru.Get("1")) // 应该输出 1
	}()

	go func() {
		defer wg.Done()
		lru.Put("3", 3) // 移除键 2
	}()

	wg.Wait()

	fmt.Println(lru.Get("2")) // 应该输出 -1
}

*/
