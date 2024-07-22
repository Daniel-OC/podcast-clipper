package download

import(
	"fmt"
	"net/url"
	"strings"
	"net/http"
  "io/ioutil"
	
	"github.com/daniel-oc/podcast-clipper/pkg/errors"

	"github.com/gocolly/colly/v2"
)

func GetPodcastDetails(urlString string) (podcastName string, episodeName string, err error) {

	podcastName, err = GetEpisodeName(urlString)
	if err != nil {
		return "", "", err
	}

	episodeName, err = GetPodcastName(urlString)
	if err != nil {
		return "", "", err
	}

	return podcastName, episodeName, nil
}

func GetEpisodeName(urlString string) (episodeName string, err error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	// decode in case of encoded values in the url name
	pathName, err := url.QueryUnescape(parsedURL.Path)
	if err != nil {
		return "", errors.NewCustomError(500, "decode", "error while decoding")
	}

	pathArray := strings.Split(pathName, "/")
	if len(pathArray) < 3 {
		return "", errors.NewCustomError(400, "invalid_path", "url path did not contain expected segments")
	}

	episodeName = pathArray[3]
	if episodeName == "" {
		return "", errors.NewCustomError(400, "empty_name", "Episode Name is empty")
	}
	
	return episodeName, nil
}

func GetPodcastName(url string) (podcastName string, err error) {
	// Scraping occurs here to get the podcast name after visiting the apple URL
	c := colly.NewCollector(
		colly.AllowedDomains("podcasts.apple.com"),
	)

	c.OnHTML(".product-header__identity.podcast-header__identity", func(element *colly.HTMLElement) {
		podcastName = element.ChildText("a")
	})

	err = c.Visit(url)
	if err != nil {
		return "",  fmt. Errorf("failed to visit URL: %w", err)
	}

	if podcastName == "" {
		return "", errors.NewScrapingError("No podcast name found")
	}

	return podcastName, nil
}

// func ConstructPodcastIndexURL(podcastName string) (URL string) {
// 	encodedPodcastName := url.QueryEscape(podcastName)
// 	URL = fmt.Sprintf("https://podcastindex.org/search?q=%s&type=all", encodedPodcastName)
// 	return URL
// }

// TODO: delete along with above functions once certain that new method gets us RSS feed
// func SearchPodcastIndexForPodcast(url string, podcastName string) (episodeURL string, err error) {
// 	podTitle := ""
// 	podLink := ""
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("podcastindex.org"),
// 	)

// 	c.OnHTML(".result-title", func(element *colly.HTMLElement) {
// 		podTitle = element.ChildText("a")
// 		podLink = element.ChildAttr("a", "href")
// 		fmt.Printf(podTitle, podLink)
// 	})
	
// 	err = c.Visit(url)
// 	if err != nil {
// 		return "", fmt. Errorf("failed to visit URL: %w", err)
// 	}

// 	if podTitle == podcastName {
// 		return podLink, nil
// 	}
	
// 	return "", fmt.Errorf("no podcast found with name: %s", podcastName)
// }

func SearchITunes(podcastName string) {
    baseURL := "https://itunes.apple.com/search"
    params := url.Values{}
    params.Add("term", podcastName)
    params.Add("entity", "podcast")

    fullURL := baseURL + "?" + params.Encode()

    resp, err := http.Get(fullURL)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    fmt.Println(string(body))
}