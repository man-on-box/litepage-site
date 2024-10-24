package data

import "fmt"

type ViewData struct {
	Homepage HomepageView
}

type HomepageView struct {
	Title      string        `json:"title"`
	Heading    string        `json:"heading"`
	Subheading string        `json:"subheading"`
	Ctas       []ctaItem     `json:"ctas"`
	Taglines   []taglineItem `json:"taglines"`
}

type ctaItem struct {
	Label    string `json:"label"`
	Url      string `json:"url"`
	Type     string `json:"type"`
	External bool   `json:"external"`
}

type taglineItem struct {
	Icon string `json:"icon"`
	Text string `json:"text"`
}

func getHomepageView() HomepageView {
	data := HomepageView{}
	parseJSONFile("./content/homepage.json", &data)

	fmt.Println(data)
	return data
}
