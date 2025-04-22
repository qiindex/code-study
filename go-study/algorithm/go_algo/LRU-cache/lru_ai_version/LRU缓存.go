package lru_ai_version

import (
	"container/list"
	"sync"
)

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

// LRUCache 结构体
type LRUCache struct {
	capacity int                      // 缓存容量
	cache    map[string]*list.Element // 哈希表，存储键值对
	list     *list.List               // 双向链表，维护访问顺序
	mu       sync.Mutex               // 互斥锁，保证线程安全
}

// entry 是存储在链表中的元素类型
type entry struct {
	key   string
	value interface{}
}

// NewLRUCache 创建一个新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// Get 获取缓存中的值
func (l *LRUCache) Get(key string) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if elem, ok := l.cache[key]; ok {
		// 移动到链表头部表示最近使用
		l.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// Put 添加或更新缓存
func (l *LRUCache) Put(key string, value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 如果键已存在，更新值并移动到头部
	if elem, ok := l.cache[key]; ok {
		elem.Value.(*entry).value = value
		l.list.MoveToFront(elem)
		return
	}

	// 如果缓存已满，淘汰最近最少使用的元素
	if l.list.Len() >= l.capacity {
		// 获取尾部元素
		tail := l.list.Back()
		if tail != nil {
			// 从哈希表中删除
			delete(l.cache, tail.Value.(*entry).key)
			// 从链表中删除
			l.list.Remove(tail)
		}
	}

	// 添加新元素到链表头部
	elem := l.list.PushFront(&entry{key, value})
	l.cache[key] = elem
}

// 底下这两个没啥用
//
//// Len 返回当前缓存中的元素数量
//func (l *LRUCache) Len() int {
//	l.mu.Lock()
//	defer l.mu.Unlock()
//	return l.list.Len()
//}
//
//// Clear 清空缓存
//func (l *LRUCache) Clear() {
//	l.mu.Lock()
//	defer l.mu.Unlock()
//	l.cache = make(map[string]*list.Element)
//	l.list = list.New()
//}
