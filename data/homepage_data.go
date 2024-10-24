package data

import "fmt"

type ViewData struct {
	Homepage HomepageView
}

type HomepageView struct {
	Title      string    `json:"title"`
	Heading    string    `json:"heading"`
	Subheading string    `json:"subheading"`
	Cta        []ctaItem `json:"cta"`
}

type ctaItem struct {
	Label    string `json:"label"`
	Url      string `json:"url"`
	Type     string `json:"type"`
	External bool   `json:"external"`
}

func getHomepageView() HomepageView {
	data := HomepageView{}
	parseJSONFile("./content/homepage.json", &data)

	fmt.Println(data)
	return data
}
