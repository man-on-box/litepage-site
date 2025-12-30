package main

import (
	"io"
	"log"

	"github.com/man-on-box/litepage"
	"github.com/man-on-box/litepage-site/data"
	"github.com/man-on-box/litepage-site/view"
)

func main() {
	domain := "litepage.dev"
	lp, err := litepage.New(domain)
	if err != nil {
		log.Fatalf("Could not create litepage app: %v", err)
	}

	d := data.New(domain)
	vh := view.NewViewHandler(d.PageData)

	lp.Page("/index.html", func(w io.Writer) {
		vh.Index.Render(w, d.Homepage)
	})

	for _, docPage := range d.DocPages {
		lp.Page(docPage.Path+".html", func(w io.Writer) {
			vh.Docs.Render(w, docPage)
		})
	}

	lp.BuildOrServe()
}
