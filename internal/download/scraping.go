package download

import(
	"net/url"
	"strings"
	"github.com/daniel-oc/podcast-clipper/pkg/errors"
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
	return "", nil
}

