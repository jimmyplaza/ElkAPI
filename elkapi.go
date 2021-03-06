package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

func timeoutDialer(cTimeout, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func ElkInput(elkurl, index, es_type string, obj interface{}) {
	var url string
	url = elkurl + index + "/" + es_type + "/"
	var myClient = &http.Client{
		Transport: &http.Transport{
			Dial: timeoutDialer(time.Duration(10)*time.Second,
				time.Duration(10)*time.Second),
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
	out, _ := json.Marshal(obj)
	outReader := bytes.NewReader(out)
	res, err := myClient.Post(url, "application/x-www-form-urlencoded", outReader)
	if err != nil {
		fmt.Printf("\nELK Error:%s\n", err)
		res.Body.Close()
		return
	}
	if res.StatusCode == 200 || res.StatusCode == 201 {
		fmt.Printf("\nELK input success, source:%v\n", string(out))
	} else {
		fmt.Printf("\nELK Error code: %d,url:%s\n", res.StatusCode, url)
	}
	res.Body.Close()
	//err = json.Unmarshal([]byte(res.Body), &obj)
	/*
		contents, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Read Body Error:%s\n", err)
			res.Body.Close()
		}
		var jobj interface{}
		err = json.Unmarshal(contents, &jobj)
		if err != nil {
			fmt.Printf("Unmarshall JSON Error:%s => %s\n", url, err)
		}
	*/
}

func ElkGet(elkurl, index, es_type string, int_id int) {
	id := strconv.Itoa(int_id)
	url := elkurl + index + "/" + es_type + "/" + id + "/"
	var myClient = &http.Client{
		Transport: &http.Transport{
			Dial: timeoutDialer(time.Duration(10)*time.Second,
				time.Duration(10)*time.Second),
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
	response, err := myClient.Get(url)
	if err != nil {
		log.Fatalf("http.Get => %v", err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("\n%v\n\n", string(body))
}

func ElkGetAll(elkurl, index, es_type string) {
	url := elkurl + index + "/" + es_type + "/_search?pretty"
	var myClient = &http.Client{
		Transport: &http.Transport{
			Dial: timeoutDialer(time.Duration(10)*time.Second,
				time.Duration(10)*time.Second),
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
	response, err := myClient.Get(url)
	if err != nil {
		log.Fatalf("http.Get => %v", err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("\n%v\n\n", string(body))
}
