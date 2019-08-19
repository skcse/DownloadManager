package download

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var folderPath ="/tmp"

type Download interface {
	startDownload(http.ResponseWriter,*http.Request,[]string)
}

type linkDownload struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}

type downloadId struct {
	Id string
}

func callDownload(down Download,writer http.ResponseWriter,request *http.Request,urls []string) {
	down.startDownload(writer,request,urls)
}

func DownloadFunc(writer http.ResponseWriter, request *http.Request)  {
	var linkBody linkDownload
	r,e1:=ioutil.ReadAll(request.Body)
	if e1!=nil{
		writer.WriteHeader(401)
		fmt.Fprintln(writer,"Cannot read payload")
	}

	e2:=json.Unmarshal(r,&linkBody)
	if e2!=nil{
		writer.WriteHeader(401)
		fmt.Fprintln(writer,"Cannot parse payload")
	}

	if linkBody.Type =="serial"{
		serialObj:=serial{Urls:linkBody.Urls}
		callDownload(serialObj,writer,request,linkBody.Urls)

	}else if linkBody.Type =="concurrent"{
		concurrentObj:=concurrent{Urls:linkBody.Urls}
		callDownload(concurrentObj,writer,request,linkBody.Urls)

	}else{
		writer.WriteHeader(401)
		fmt.Fprintln(writer,"Download type doesnt exit:",linkBody.Type)
	}
}

func downloadFile(url string,id string) error {

	filepath:=folderPath+"/"+ id
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

func generateId() string  {
	id:=uuid.New()
	return id.String()
}

