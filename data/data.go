package data

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type viewData struct {
	Homepage homepageView
	DocPages *[]*docPage
	PageData *PageData
}

type PageData struct {
	Links        Links
	DocsSections *[]docsSection
}

func New() *viewData {
	docs := parseDocs()
	pd := &PageData{
		Links:        getLinksView(),
		DocsSections: docs.Sections,
	}
	d := &viewData{
		Homepage: getHomepageView(),
		DocPages: docs.Pages,
		PageData: pd,
	}
	return d
}

func parseJSONFile(file string, data any) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer f.Close()

	contents, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	err = json.Unmarshal(contents, &data)
	if err != nil {
		log.Fatalf("Could not parse JSON: %v", err)
	}
}
