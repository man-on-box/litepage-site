package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type viewData struct {
	Homepage homepageView
	DocPages *[]docPage
	PageData *PageData
}

type PageData struct {
	Links       Links
	DocContents *[]docsSection
}

func New() *viewData {
	docs := parseDocs()
	pd := &PageData{
		Links:       getLinksView(),
		DocContents: docs.Contents,
	}
	for _, t := range *docs.Contents {
		fmt.Println(t.Label)
		fmt.Println(len(t.Items))

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
