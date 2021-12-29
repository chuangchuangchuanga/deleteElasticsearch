package main

import (
	"context"
	"flag"
	"fmt"

	e "github.com/olivere/elastic/v7"
)

func main() {

	index := flag.String("index", "", "索引名")
	url := flag.String("url", "", "es地址")
	flag.Parse()

	client, err := e.NewClient(e.SetSniff(false), e.SetURL(*url))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	query := e.Query(e.NewRangeQuery("@timestamp").Gte("now-100d").Lte("now-7d"))

	fmt.Print(query)

	resule, err := client.DeleteByQuery().Index(*index).Query(query).Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	} else {
		fmt.Println(resule)
	}

}
