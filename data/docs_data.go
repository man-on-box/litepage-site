package data

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
)

type docsView struct {
	Sections []docsSection
	Pages    []*docPage
}

type docsSection struct {
	Label string
	Items []sectionItem
}

type sectionItem struct {
	Label string
	Path  string
}

type docPage struct {
	Path        string
	Title       string
	Description string
	Section     string
	Content     template.HTML
	Next        *docPage
	Previous    *docPage
}

type docFrontmatter struct {
	Slug        string `yaml:"slug"`
	Title       string `yaml:"title"`
	Section     string `yaml:"section"`
	Description string `yaml:"description"`
}

const docsDir = "content/docs/"

func parseDocs() docsView {
	sections := []docsSection{}
	pages := []*docPage{}
	entries, err := os.ReadDir(docsDir)
	if err != nil {
		log.Fatalf("Error reading docs directory: %v", err)
	}

	// We'll assume the order of files are also by section.
	// This way I don't need to create an ordered map
	for index, entry := range entries {
		entryPath := docsDir + entry.Name()
		fm := &docFrontmatter{}
		html := parseMdFile(entryPath, fm)

		docsPath := "/docs/" + fm.Slug

		isNewSection := len(sections) == 0 || sections[len(sections)-1].Label != fm.Section
		if isNewSection {
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

		page := &docPage{
			Path:        docsPath,
			Title:       fm.Title,
			Description: fm.Description,
			Section:     fm.Section,
			Content:     html,
		}

		if index > 0 {
			previousPage := pages[index-1]
			page.Previous = previousPage
			previousPage.Next = page

		}

		pages = append(pages, page)
	}

	return docsView{
		Sections: sections,
		Pages:    pages,
	}
}

var mdParser = goldmark.New(
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
	goldmark.WithExtensions(
		&frontmatter.Extender{},
		highlighting.NewHighlighting(
			highlighting.WithStyle("vulcan"),
		),
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
