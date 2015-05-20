package feedReader

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type FeedXML struct {
	Title string `xml:"channel>title"`
	Episodes []EpisodeXML `xml:"channel>item"`
}

type EpisodeXML struct {
	Data string `xml:"pubDate"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	Title string `xml:"title"`
}

func GetEpisodesfromFeed(feedUrl string) ([]EpisodeXML, error) {
	if resp, err := http.Get(feedUrl); err == nil {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var feed FeedXML
			if err = xml.Unmarshal([]byte(body), &feed) ; err == nil {
				return feed.Episodes, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
