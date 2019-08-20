package download

import (
	"encoding/json"
	"github.com/skcse/DownloadManager/source/status"
	"net/http"
	"time"
)

type concurrent struct {
	Urls []string
}

func (c concurrent) startDownload(writer http.ResponseWriter,request *http.Request,urls []string)  {
	startTime:=time.Now()
	mapFiles:=make(map[string]string)
	finalMapFiles:=make(map[string]string)
	for _,url:=range urls{
		name := generateId()
		mapFiles[url]=name
	}
	const limitWorker =2
	reqChan:=make(chan string)

	for i:=0;i<limitWorker;i++{
		go worker(reqChan,mapFiles,finalMapFiles,len(urls))
	}

	id:= generateId()
	idData:= downloadId{Id: id}
	writer.Header().Set("Content-Type", "application/json")
	js,_:=json.Marshal(idData)

	status.Mp[id]= status.Status{id,startTime,time.Now(),"QUEUED","CONCURRENT",mapFiles}

	go func() {
		for _, url := range urls {
			reqChan <- url
		}
		status.Mp[id]= status.Status{id,startTime,time.Now(),"SUCCESSFUL","CONCURRENT",mapFiles}
		return
	}()

	writer.Write(js)
}

func worker(reqChan chan string,mapFiles map[string]string,finalMapFiles map[string]string,urlsCount int)  {
	for {
		select {
			case url,ok := <-reqChan:
				name:=mapFiles[url]
				_ = downloadFile(url, name)
				finalMapFiles[url] = name
				if !ok {
					return
				}
			}
			if urlsCount == len(finalMapFiles){
				close(reqChan)
				return
			}
		}
}

