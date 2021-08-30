package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	e "github.com/olivere/elastic/v7"
)

func main() {

	index := flag.String("index", "", "索引名")
	url := flag.String("url", "", "es地址")
	deletetime := flag.String("deletetime", "-240h", "删除前多少天的数据,默认10天")
	flag.Parse()

	client, err := e.NewClient(e.SetSniff(false), e.SetURL(*url))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	now := time.Now()
	d, _ := time.ParseDuration(*deletetime)
	deleteTime := now.Add(d).Format("2006-01-02")

	query := e.Query(e.NewRangeQuery("timestamp").Lte(deleteTime))

	fmt.Print(query)

	resule, err := client.DeleteByQuery().Index(*index).Query(query).Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	} else {
		fmt.Println(resule)
	}

}
