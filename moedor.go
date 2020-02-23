package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://www.deviante.com.br/podcasts/scicast/feed/")
	fmt.Println(fmt.Sprintf("%v", feed))

	// for _, item := range feed.Items {
	// 	fmt.Println(item.ITunesExt.Duration)
	// 	fmt.Println(item.PublishedParsed.Unix())
	// }
}
