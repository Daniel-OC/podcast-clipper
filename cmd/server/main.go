package main

import (
	"fmt"
	"github.com/daniel-oc/podcast-clipper/internal/download"
)

func main() {
	testURL:= "bill simmons"

	podcastName := download.ConstructPodcastIndexURL(testURL)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }

	fmt.Printf("Episode Name: %s\n", podcastName)
}