package main

import (
	"go-study/algorithm/demo_design"
	"go-study/algorithm/go_algo"
	"go-study/business/list_info"
	"go-study/business/string_info"
)

func demo1() {
	println("222")
	// is the demo
	println(list_info.SortInfo(2))
	name := string_info.UpperName("New world")
	println(name)
	var a int
	var b int = 2
	println(a + b)
}

func main() {
	// demo
	demo1()
	demo_design.Demo()
	go_algo.CompareString()
	//go_algo.Main1()
	//LRU缓存_手写版_加锁类型.NewLruCache(2)
	//lru_cache.Constructor(2)
}
