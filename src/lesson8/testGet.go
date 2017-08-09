package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//get tls加密链接
	var tr = &http.Transport{
		//TLSClientConfig: &tls.Config{InsecureSkipVerify: true,Renegotiation: tls.RenegotiateFreelyAsClient,},
		TLSClientConfig: &tls.Config{Renegotiation: tls.RenegotiateFreelyAsClient},
	}
	var client = &http.Client{Transport: tr}
	url := "https://restapi.cdn.azure.cn/subscriptions/ee57ca15-7f92-4a93-ac70-f0888a2fb0e3/endpoints?apiVersion=1"
	resp, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
}
