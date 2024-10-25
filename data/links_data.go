package data

type Links struct {
	Urls map[string]url `json:"urls"`
	Nav  []navItem      `json:"nav"`
}

type url struct {
	Url    string `json:"url"`
	Target string `json:"target"`
}

type navItem struct {
	Label string `json:"label"`
	Key   string `json:"urlKey"`
}

func getLinksView() Links {
	links := Links{}
	parseJSONFile("./content/links.json", &links)
	return links
}
