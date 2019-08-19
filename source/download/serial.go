package download

import (
	"encoding/json"
	"fmt"
	"github.com/skcse/DownloadManager/source/status"
	"net/http"
	"sync"
	"time"
)

func serialDownload(writer http.ResponseWriter,request *http.Request,urls []string,starttime time.Time)  {
	mapFiles:=make(map[string]string)
	for _,url:= range urls{
		name:= generateId()
		_= downloadFile(url,name)
		mapFiles[url]=name
	}
	var wg sync.WaitGroup
	wg.Add(1)
	id:= generateId()
	idData:= downloadId{Id: id}
	writer.Header().Set("Content-Type", "application/json")
	js,_:=json.Marshal(idData)

	status.Mp[id]= status.Status{id,starttime,time.Now(),"SUCCESSFUL","SERIAL",mapFiles}

	go writer.Write(js)
	go call(&wg)
	fmt.Fprintln(writer,"hi there!!")
	wg.Wait()
}

func call(wg *sync.WaitGroup)  {
	time.Sleep(10*time.Second)
	wg.Done()

}