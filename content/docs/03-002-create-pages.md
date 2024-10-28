---
slug: create-pages
title: Create pages
section: Methods and usage
description: Learn about Litepage methods and their usage
---

# Creating pages

Create new pages with the `Page` method by passing in the relative filename that will be used when building the site, such as `/index.html` or nested pages like `/articles/new-recipes.html`.

```go
lp, _ := litepage.New("hello-world.com")

err := lp.Page("/index.html", func (w io.Writer) {
		html := "<h1>Hello, World!</h1>"
		t := template.Must(template.New("helloWorld").Parse(html))
	    t.Execute(w, nil)
})
```

_**Note:** Paths must start with a forward slash `/`, include a file extension and be a valid filepath._

Here you also pass a function that receives the standard `io.Writer` interface. Write your templates to this interface to generate your pages with your content.

This means you can use any html templating library that supports this interface, such as the Go standard [html/template](https://pkg.go.dev/html/template) package or custom templating packages like [templ](https://templ.guide/).

An error is returned if the file path is not valid, or if you have already created a page with the same file path previously.

## Flexible Data Sources and Template Handling

Note how Litepage has no opinion on where your data is coming from, or how you build your templates. You can add as many pages as you need, and call this function however you like. The source of your data can be from some database, CMS, external API or files in your project.