package main

import (
	"github.com/olivere/elastic/v7",
	"fmt"
)


func main(){

	client, err := elastic.NewClient(elastic.SetURL("http://192.168.2.10:9201"))
	if err != nil {
		// Handle error
		fmt.printf("连接错误")
	}




}