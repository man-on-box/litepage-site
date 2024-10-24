package view

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/man-on-box/litepage-site/data"
)

type ViewHandler struct {
	Index *view
}

func NewViewHandler() *ViewHandler {
	return &ViewHandler{
		Index: newView("base", "view/index.html"),
	}
}

type viewData struct {
	Data any
}

type view struct {
	Layout   string
	Template *template.Template
}

var tmpl = template.New("").Funcs(template.FuncMap{
	"version": func() string {
		return time.Now().Format("01021504")
	},
	"inlineSVG": inlineSVG,
})

func newView(layout string, files ...string) *view {
	files = append(layoutFiles(), files...)
	t, err := tmpl.ParseFiles(files...)
	if err != nil {
		log.Fatalf("error while parsing files: %v", err)
	}

	return &view{
		Layout:   layout,
		Template: t,
	}
}

func (v *view) Render(w io.Writer, data data.HomepageView) {
	vd := viewData{
		Data: data,
	}

	err := v.Template.ExecuteTemplate(w, v.Layout, vd)
	if err != nil {
		log.Fatalf("error while rendering template: %v", err)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob("view/layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func inlineSVG(filename string, className ...string) template.HTML {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("could not open svg %s: %v", filename, err)
	}

	svg := string(content)

	if len(className) > 0 {
		svg = strings.Replace(svg, "<svg", `<svg class="`+className[0]+`"`, 1)
	}

	return template.HTML(svg)
}
