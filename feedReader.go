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
	Date string `xml:"pubDate" json:"date"`
	Description string `xml:"description" json:"description"`
	Itunes Itunes `xml:"duration" json:"info"`
	Link string `xml:"link" json:"link"`
	Title string `xml:"title" json:"title"`
	Content string `xml"enclosure url,attr" json:"content"`
	Enclosure Enclosure `xml:"enclosure" json:"enclosure"`
	Dc Dc `xml:"creator" json:"author"`
}

type Enclosure struct {
	Content string `xml:"url,attr" json:"url"`
	Type string `xml:"type,attr" json:"type"`
}

type Itunes struct {
	Duration string `xml:",chardata" json:"duration"`
}

type Media struct {
	Thumbnail string `xml:"thumbnail,attr" json:"thumbnail"`
}

type Dc struct {
	Creator string `xml:",chardata" json:"name"`
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
