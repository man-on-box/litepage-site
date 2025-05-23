package view

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/man-on-box/litepage-site/data"
)

type ViewHandler struct {
	Index *view
	Docs  *view
}

func NewViewHandler(pd *data.PageData) *ViewHandler {
	return &ViewHandler{
		Index: newView(pd, "base", "view/index.html"),
		Docs:  newView(pd, "base-docs", "view/doc-page.html"),
	}
}

type viewData struct {
	PageData *data.PageData
	Data     any
}

type view struct {
	Layout   string
	Template *template.Template
	PageData *data.PageData
}

func newView(pd *data.PageData, layout string, files ...string) *view {
	files = append(layoutFiles(), files...)

	tmpl := template.New("").Funcs(template.FuncMap{
		"version": func() string {
			return fmt.Sprintf("%d", time.Now().Unix())
		},
		"currentYear": func() string {
			return strconv.Itoa(time.Now().Year())
		},
		"inlineSVG": inlineSVG,
	})

	t, err := tmpl.ParseFiles(files...)
	if err != nil {
		log.Printf("error while parsing files: %v", err)
	}

	return &view{
		Layout:   layout,
		Template: t,
		PageData: pd,
	}
}

func (v *view) Render(w io.Writer, data any) {
	vd := viewData{
		Data:     data,
		PageData: v.PageData,
	}

	err := v.Template.ExecuteTemplate(w, v.Layout, vd)
	if err != nil {
		log.Printf("error while rendering template: %v", err)
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
		classes := strings.Join(className, " ")
		svg = strings.Replace(svg, "<svg", `<svg class="`+classes+`"`, 1)
	}

	return template.HTML(svg)
}
