package download

import (
	"net/http"
	"sync"
)

type concurrent struct {
	Urls []string
}

func (c concurrent) startDownload(writer http.ResponseWriter,request *http.Request,urls []string)  {
	var wg sync.WaitGroup
	wg.Add(1)
	//req:=make(chan string)
	//
	//go gatherResults(posMap, resChan, reqChan,len(words), &wg)
}
