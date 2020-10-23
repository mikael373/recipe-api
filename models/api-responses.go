package models

type PuppyResponse struct {
	Results []struct {
		Title       string `json:"title"`
		Href        string `json:"href"`
		Ingredients string `json:"ingredients"`
		Thumbnail   string `json:"thumbnail"`
	} `json:"results"`
}

type GiphyResponse struct {
	Data []struct {
		Images struct {
			Original struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
}
