package main

import (
	"fmt"
	"github.com/daniel-oc/podcast-clipper/internal/download"
)

func main() {
	testURL:= "https://podcasts.apple.com/us/podcast/episode-57-the-great-orphan-siege/id1539317046?i=1000554129416"

	podcastName, err := download.GetPodcastName(testURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Episode Name: %s\n", podcastName)
}