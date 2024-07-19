package download

import(
	"net/url"
	"strings"
	"github.com/daniel-oc/podcast-clipper/pkg/errors"
)

func ValidateApplePodcastURL(url string) error {
	parsedURL, err := url.Parse(url)
	if err != nil {
		return errors.NewInvalidURLError("failed to parse url")
	}

	if !strings.HasSuffix(parsedURL.Hostname(), "apple.com") {
		return errors.NewInvalidURLError("Not an Apple Podcasts URL")
	}

	return nil
}