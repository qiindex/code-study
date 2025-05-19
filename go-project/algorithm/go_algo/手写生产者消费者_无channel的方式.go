package go_algo

import (
	"sync"
)

type Buffer struct {
	data  []int
	lock  sync.Mutex
	size  int
	count int
	cond  *sync.Cond
}

func NewBuffer(size int) *Buffer {
	b := &Buffer{
		data: make([]int, 0),
		size: size,
		//count: ,
	}
	b.cond = sync.NewCond(&b.lock)
	return b
}

// pr
func (b *Buffer) produce(item int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	for b.count == b.size {
		b.cond.Wait()
	}
	//.....添加元素
	//noti
	b.cond.Signal()
}

// cons
func (b *Buffer) Consume() int {
	b.lock.Lock()
	defer b.lock.Unlock()

	for b.count == 0 {
		b.cond.Wait()
	}
	//消费。。。
	b.cond.Signal()

	//提交offset。或者消费完成的结果
	return 0
}
