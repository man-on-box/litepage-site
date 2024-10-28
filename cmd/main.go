package main

import (
	"io"
	"log"

	"github.com/man-on-box/litepage"
	"github.com/man-on-box/litepage-site/data"
	"github.com/man-on-box/litepage-site/view"
)

func main() {
	lp, err := litepage.New("litepage.dev")
	if err != nil {
		log.Fatalf("Could not create litepage app: %v", err)
	}

	d := data.New()
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
