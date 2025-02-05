---
slug: markdown
title: Markdown
section: Common recipes
description: How to get setup using Markdown in your Litepage site
---

# Using Markdown

Markdown is a common syntax to create formatted documents in a plain text format. This works really well when you want to encapsulate formatted content within your project (**tip**, the docs you are reading right now are all [formatted with markdown](https://github.com/man-on-box/litepage-site/tree/main/content/docs)).

Golang does not include any Markdown parser in its standard library, so we can reach for a package to do this for us.

There are a few packages available but we can recommend [goldmark](https://github.com/yuin/goldmark) because of its stability and extendability (if you want to use extra features like [frontmatter](https://github.com/abhinav/goldmark-frontmatter) for example).

With this, you can start to parse Markdown and include this into your HTML templates.

## Full example

_(We will explain the example in more detail below)_

```go
package main

import (
	"html/template"

	"github.com/man-on-box/litepage"
	"github.com/yuin/goldmark"
)

// You would grab this from a file somewhere
const exampleMarkdown = `
## Hello world

You are looking **great** today.
`

// And probably you would grab your template from a file somewhere
const exampleTemplate = `
<div>
	<h1>A fancy page with content from Markdown</h1>
	<article>
		{{ . }}
	</article>
</div>
`

func main() {
  	lp, _ := litepage.New("litepage.dev")

	lp.Page("/temp.html", func(w io.Writer) {

		t, err := template.New("pageWithMarkdown").Parse(exampleTemplate)
		if err != nil {
			log.Fatalf("error parsing template: %v", err)
		}

		var buf bytes.Buffer
		err = goldmark.Convert([]byte(exampleMarkdown), &buf)
		if err != nil {
			log.Fatalf("error parsing markdown: %v", err)
		}

		mdHtml := template.HTML(buf.String())

		t.Execute(w, mdHtml)
	})

	lp.BuildOrServe()
}
```

First we are getting our `exampleMarkdown` and `exampleTemplate` strings. Normally you would get these from their own `.md` and `.html` files in your repository. Notice in the html template, we are passing `{{ . }}`, this is where the rendered markdown will be placed, (inside the `article` tag).

Next in `main()` we are creating a new page of `/temp.html`. Within here we are creating a new template from the example HTML. We are also parsing the markdown with **goldmark**, and creating a new `html.Template` with the parsed string.

Finally, we pass the parsed markdown HTML to the template we created with `t.Execute(w, mdHtml)`. This executes the template and writes it to the `io.Writer` that we pass in from Litepage.

**If this looks tricky** don't worry, you can abstract this logic out to its own function so you only need write it once, and it's not a lot of code. Once you have it, you now have full markdown support in your project, and you only brought in **one dependency**.

I would recommend to read through the [goldmark](https://github.com/yuin/goldmark) documentation for better understanding of features and plugins you can use.
