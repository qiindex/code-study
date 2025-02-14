package main

import (
	list_info  "go_study/business/list_info"
	stringinfo "go_study/business/stringinfo"
)

func main() {
	println("222")
	// is the demo
	println(list_info.SortInfo(2))
	name := stringinfo.UpperName("New world")
	println(name)

}
