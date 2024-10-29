package data

type homepageView struct {
	Title           string        `json:"title"`
	Heading         string        `json:"heading"`
	Subheading      string        `json:"subheading"`
	Ctas            []ctaItem     `json:"ctas"`
	Taglines        []taglineItem `json:"taglines"`
	SnippetsHeading string        `json:"snippetsHeading"`
	Snippets        []snippetItem `json:"snippets"`
}

type ctaItem struct {
	Label string `json:"label"`
	Key   string `json:"urlKey"`
}

type taglineItem struct {
	Icon string `json:"icon"`
	Text string `json:"text"`
}

type snippetItem struct {
	Heading     string `json:"heading"`
	Description string `json:"description"`
	Snippet     string `json:"snippet"`
	ColSpan     int    `json:"col-span"`
}

func getHomepageView() homepageView {
	data := homepageView{}
	parseJSONFile("./content/homepage.json", &data)
	return data
}
