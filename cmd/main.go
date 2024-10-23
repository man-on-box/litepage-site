package main

import (
	"io"
	"log"

	"github.com/man-on-box/litepage"
	"github.com/man-on-box/litepage-site/view"
)

func main() {
	lp, err := litepage.New("litepage.dev")
	if err != nil {
		log.Fatalf("Could not create litepage app: %v", err)
	}

	vh := view.NewViewHandler()

	lp.Page("/index.html", func(w io.Writer) {
		vh.Index.Render(w, nil)
	})

	lp.BuildOrServe()
}
