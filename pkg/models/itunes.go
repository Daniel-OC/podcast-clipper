package models

type PodcastResult struct {
    CollectionName string   `json:"collectionName"`
    ArtistName     string   `json:"artistName"`
    FeedURL        string   `json:"feedUrl"`
    Genres         []string `json:"genres"`
}

type FullPodcastResultResponse struct {
	ResultCount int `json:"resultCount"`
	Results []PodcastResult `json:"results"`
}