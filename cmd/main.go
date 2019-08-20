package main

import (
	"fmt"
	"github.com/skcse/DownloadManager/source/download"
	"github.com/skcse/DownloadManager/source/status"
	"log"
	"net/http"
)

func main()  {
	h:=http.NewServeMux()

	h.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET"{
			writer.WriteHeader(200)
			fmt.Fprintln(writer,"OK")
		}
	})
	h.HandleFunc("/downloads", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST"{
			download.DownloadFunc(writer,request)
		}else{
			writer.WriteHeader(404)
			fmt.Fprintln(writer,"NOT Found")
		}
	})
	h.HandleFunc("/downloads/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET"{
			status.GetStatus(writer,request)
		}else{
			writer.WriteHeader(404)
			fmt.Fprintln(writer,"NOT Found")
		}
	})
	err:=http.ListenAndServe(":8000",h)
	log.Fatal(err)
}

func home(writer http.ResponseWriter, request *http.Request)  {
	writer.WriteHeader(200)
	fmt.Fprintln(writer,"Hi there")
}
