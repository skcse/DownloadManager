package download

import (
	"encoding/json"
	"github.com/skcse/DownloadManager/source/status"
	"net/http"
	"strings"
	"time"
)

type concurrent struct {
	Urls []string
}

func (c concurrent) startDownload(writer http.ResponseWriter,request *http.Request,urls []string)  {
	startTime:=time.Now()
	mapFiles:=make(map[string]string)
	folder:=generateId()
	finalMapFiles:=make(map[string]string)
	for _,url:=range urls{
		name := generateId()
		mapFiles[url]=FolderPath + "/" +folder+ "/" + name
	}
	const limitWorker =6
	reqChan:=make(chan string)

	id:= generateId()
	idData:= downloadId{Id: id}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	js,_:=json.Marshal(idData)

	for i:=0;i<limitWorker;i++{
		go worker(reqChan,mapFiles,finalMapFiles,len(urls),id,startTime)
	}

	status.Mp[id]= status.Status{id,startTime,time.Now(),"QUEUED","CONCURRENT",mapFiles}

	go func() {
		for _, url := range urls {
			reqChan <- url
		}
		return
	}()

	writer.Write(js)
}

func worker(reqChan chan string,mapFiles map[string]string,finalMapFiles map[string]string, urlsCount int,downId string,startTime time.Time)  {
	for {
		select {
			case url,ok := <-reqChan:
				if !ok {
					return
				}
				folderPath:=mapFiles[url]
				arr:=strings.Split(folderPath,"/")
				_ = downloadFile(url,arr[3] ,arr[2])
				finalMapFiles[url] = arr[3]
			}
			if urlsCount == len(finalMapFiles){
				status.Mp[downId]= status.Status{downId,startTime,time.Now(),"SUCCESSFUL","CONCURRENT",mapFiles}
				close(reqChan)
				return
			}
	}
}