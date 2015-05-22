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
	Content string `xml:"content" json:"content"`
	Date string `xml:"pubDate" json:"date"`
	Description string `xml:"description" json:"description"`
	Itunes Itunes `xml:"duration" json:"info"`
	Link string `xml:"link" json:"link"`
	Title string `xml:"title" json:"title"`
}

type Itunes struct {
	Duration string `xml:",chardata"`
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
