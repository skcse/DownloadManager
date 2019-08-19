package download

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/google/uuid"
	"time"
)

var folderPath ="/tmp"

type linkDownload struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}

type downloadId struct {
	Id string
}

func Download(writer http.ResponseWriter, request *http.Request)  {
	var linkBody linkDownload
	r,_:=ioutil.ReadAll(request.Body)
	err:=json.Unmarshal(r,&linkBody)
	if err!=nil{
		writer.WriteHeader(401)
		fmt.Fprintln(writer,"Cannot parse payload")
	}
	if linkBody.Type =="serial"{
		startTime:=time.Now()
		serialDownload(writer,request,linkBody.Urls,startTime)
	}else if linkBody.Type =="concurrent"{
		concurrent(writer,request,linkBody.Urls)
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

