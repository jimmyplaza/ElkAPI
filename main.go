package main

import (
	"fmt"
	"time"
	"strings"
)

type JsonType struct {
	Status int 			`json:"status"`
	ResponseTime string `json:"responsetime"`
	Timestamp    string `json:"@timestamp"`
	Url          string `json:"url"`
}

func main() {
	var source JsonType	
	source.Status = 100 
	source.ResponseTime = "902.1ms"
	current_time := fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04"))
	current_time = strings.Replace(current_time, " ", "T", 1)
	source.Timestamp = current_time

	url := "http://g2tool.cloudapp.net:9200/"
	index := "test_idx"
	es_type := "test_tbl"
	ElkInput(url, index, es_type, source)

	//ElkGet(url, index, es_type, 14)
	//ElkGetAll(url, index, es_type)
}

