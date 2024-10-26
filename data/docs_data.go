package data

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

// Main return. Contents are to be used in every docs page template
// Pages are to iterate over to create every docs page
type docsView struct {
	Contents *[]docsSection
	Pages    *[]docPage
}

// Contains docs sections and their items
type docsSection struct {
	Label string
	Items []sectionItem
}

// Section item, used to create a tags in contents pane
type sectionItem struct {
	Label string
	Path  string
}

// Actual docs page, includes path to create and contents
type docPage struct {
	Path        string
	Title       string
	Description string
	Content     template.HTML
}

type docFrontmatter struct {
	Slug        string `yaml:"slug"`
	Title       string `yaml:"title"`
	Section     string `yaml:"section"`
	Description string `yaml:"description"`
}

func parseDocs() docsView {
	sections := []docsSection{}
	pages := []docPage{}
	entries, err := os.ReadDir("content/docs")
	if err != nil {
		log.Fatalf("Error reading docs directory: %v", err)
	}

	// We'll assume the order of docs are also by section.
	// This way I don't need to create an ordered map
	for _, entry := range entries {
		entryPath := "content/docs/" + entry.Name()
		fm := &docFrontmatter{}
		html := parseMdFile(entryPath, fm)

		docsPath := "/docs/" + fm.Slug

		// Build content sections here...
		if len(sections) == 0 || sections[len(sections)-1].Label != fm.Section {
			s := docsSection{
				Label: fm.Section,
				Items: []sectionItem{
					{Label: fm.Title, Path: docsPath},
				},
			}
			sections = append(sections, s)
		} else {
			s := &sections[len(sections)-1]
			s.Items = append(s.Items, sectionItem{
				Label: fm.Title, Path: docsPath,
			})
		}

		// Build docs list here
		pages = append(pages, docPage{
			Path:        docsPath + ".html",
			Title:       fm.Title,
			Description: fm.Description,
			Content:     html,
		})
	}

	return docsView{
		Contents: &sections,
		Pages:    &pages,
	}
}

var mdParser = goldmark.New(
	goldmark.WithExtensions(
		&frontmatter.Extender{},
	),
)

func parseMdFile(file string, metadata any) template.HTML {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading markdown file: %v", err)
	}
	var buf bytes.Buffer
	ctx := parser.NewContext()
	if err := mdParser.Convert(f, &buf, parser.WithContext(ctx)); err != nil {
		log.Fatalf("Error converting markdown to HTML: %v", err)
	}
	if metadata != nil {
		d := frontmatter.Get(ctx)
		if err := d.Decode(metadata); err != nil {
			log.Fatalf("Error decoding frontmatter data: %v", err)
		}
	}
	return template.HTML(buf.String())
}
