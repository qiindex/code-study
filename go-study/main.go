package main

import (
	 "go-study/business/list_info"
	 "go-study/business/string_info"
)

func main() {
	demo1()
	demo2()
}
func demo1(){
	println("222")
	// is the demo
	println(list_info.SortInfo(2))
	name := string_info.UpperName("New world")
	println(name)
	var a int 
	var  b  int = 2
	println(a+ b)
}

func demo2(){

}
