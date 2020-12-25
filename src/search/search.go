package search

import (
	"log"
	"sync"
)
//注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

func Run(searcjTerm string)  {
	//获取需要搜索的数据源列表
	feeds,err:=RetrieveFeeds()
	if err!=nil {
		log.Fatal(err)
	}

	//创建一个无缓冲的通道，接受匹配后的结果
	results:=make(chan *Result)

	//构造一个waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	
}