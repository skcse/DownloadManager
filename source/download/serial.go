package download

import (
	"encoding/json"
	"github.com/skcse/DownloadManager/source/status"
	"net/http"
	"time"
)

type serial struct {
	Urls []string
}

func (s serial) startDownload(writer http.ResponseWriter,request *http.Request,urls []string)  {

	startTime:=time.Now()
	mapFiles:=make(map[string]string)
	folder:=generateId()
	for _,url:= range urls{
		name:= generateId()
		_=downloadFile(url,name,folder)
		mapFiles[url]=FolderPath + "/" + folder + "/" + name
	}
	id:= generateId()
	idData:= downloadId{Id: id}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	js,_:=json.Marshal(idData)

	status.Mp[id]= status.Status{id,startTime,time.Now(),"SUCCESSFUL","SERIAL",mapFiles}

	writer.Write(js)
}


