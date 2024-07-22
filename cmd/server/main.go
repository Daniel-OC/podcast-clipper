package main

import (
	// "fmt"
	"github.com/daniel-oc/podcast-clipper/internal/download"
)

func main() {
	// testURL:= "https://podcastindex.org/search?q=bill%26simmons&type=all"

	download.SearchITunes("the bill simmons podcast")
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }

	// fmt.Printf("Episode Name: %s\n", podcastName)
}