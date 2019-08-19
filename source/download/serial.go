package download

import (
	"encoding/json"
	"net/http"
	"github.com/skcse/DownloadManager/status"
	"time"

)

func serialDownload(writer http.ResponseWriter,request *http.Request,urls []string,starttime time.Time)  {
	mapFiles:=make(map[string]string)
	for _,url:= range urls{
		name:=generateId()
		_=downloadFile(url,name)
		mapFiles[url]=name
	}
	id:=generateId()
	idData:=downloadId{Id:id}
	writer.Header().Set("Content-Type", "application/json")
	js,_:=json.Marshal(idData)

	status.Mp[id]=status.Status{id,starttime,time.Now(),"SUCCESSFUL","SERIAL",mapFiles}

	writer.Write(js)
}
