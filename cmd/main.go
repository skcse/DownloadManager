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
			writer.WriteHeader(200)
			fmt.Fprintln(writer,"OK")
	})
	h.HandleFunc("/", home)
	h.HandleFunc("/downloads", download.DownloadFunc)
	h.HandleFunc("/getStatus/", status.GetStatus)
	err:=http.ListenAndServe(":8000",h)
	log.Fatal(err)
}

func home(writer http.ResponseWriter, request *http.Request)  {
	writer.WriteHeader(200)
	fmt.Fprintln(writer,"Hi there")
}
