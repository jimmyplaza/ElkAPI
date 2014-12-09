package main

type JsonType struct {
	Status int 		`json:"status"`
	ResponseTime int    `json:"responsetime"`
	Timestamp    string `json:"@timestamp"`
	Url          string `json:"url"`
}

func main() {
	var source JsonType	
	source.Status = 99
	source.ResponseTime = "902.1ms"
	current_time := fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04"))
	current_time = strings.Replace(current_time, " ", "T", 1)
	source.Timestamp := current_time

	url := "http://g2tool.cloudapp.net:9200/"
	index := "test_idx"
	type := "test_tbl"
	ElkInput(url, index, type, source)

	//ElkGet(index, type, 14)
	//ElkGetAll(index, type)
}

