package feedReader


func GetEpisodesfromFeed(feedUrl string) {
	if resp, err := http.Get(feedUrl); err == nil {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			log.Println(body)
		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
}
