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
	deletetime := flag.String("deletetime", "2021-08-31", "删除指定日期的数据")
	flag.Parse()

	client, err := e.NewClient(e.SetSniff(false), e.SetURL(*url))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	query := e.Query(e.NewMatchQuery("timestamp", deletetime))

	fmt.Print(query)

	resule, err := client.DeleteByQuery().Index(*index).Query(query).Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	} else {
		fmt.Println(resule)
	}

}
