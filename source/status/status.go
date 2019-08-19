package status

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

//import (
//	"database/sql",
//	"github.com/go-sql-driver/mysql"
//)

type Status struct {
	Id           string   `json:"id"`
	StartTime    time.Time   `json:"start_time"`
	EndTime      time.Time   `json:"end_time"`
	Status       string   `json:"status"`
	DownloadType string   `json:"download_type"`
	Files        map[string]string `json:"files"`
}

var Mp=make(map[string]Status)

func GetStatus(w http.ResponseWriter, r *http.Request)  {
	u, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	id:= u.RawQuery
	status:=Mp[id]
	fmt.Fprintln(w,status)
	//mp:=Mp[id]
	//js,_:=json.Marshal(mp)

}
