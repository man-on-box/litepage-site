---
slug: preview-site
title: Preview site
section: Methods and usage
description: Learn about Litepage methods and their usage
---

# Previewing your site

So we know how to build, but what about while developing your site? Ideally instead of writing files to `/dist` everytime you make a change, it would be nice to be able to see your site hosted locally on your machine for a better developer experience.

For this you can use the `Serve` method which serves your static site locally, instead of outputting it to your dist folder.

```go
lp, _ := litepage.New("hello-world.com")

// ... add all your pages

err := lp.Serve("3000")
```

This will start a web server at http://localhost:3000 to preview your site.

## Live reloading

Litepage does not include any live reloading out of the box. There are live reloading tools out there already to rebuild your project on changes. If you are not using a live reloading tool already, I can recommend [Air](https://github.com/air-verse/air), which also offers a [proxy server](https://github.com/air-verse/air/pull/512) to automatically refresh your page when you save changes.

While it is not as instant as hot reloading a JS project with [Vite](https://vite.dev/), it is fast enough to offer a nice development experience.

## Smart routing

Routing is handled for you to match those of static hosting providers. Take the following example site with three pages:

```go
lp, _ := litepage.New("hello-world.com")

lp.Page("/index.html", func (w io.Writer) {
		html := "<h1>Hello, World!</h1>"
		t := template.Must(template.New("helloWorld").Parse(html))
	    t.Execute(w, nil)
})

lp.Page("/nested/index.html", func (w io.Writer) {
		html := "<h1>Nested index page</h1>"
		t := template.Must(template.New("nestedIndex").Parse(html))
	    t.Execute(w, nil)
})

lp.Page("/foo.html", func (w io.Writer) {
		html := "<h1>Foo page</h1>"
		t := template.Must(template.New("fooPage").Parse(html))
	    t.Execute(w, nil)
})

err := lp.Serve("3000")
```

You can reach these pages with any of the following paths.

**`/index.html` page**:

- http://localhost:3000
- http://localhost:3000/
- http://localhost:3000/index
- http://localhost:3000/index.html

**`/nested/index.html` page**:

- http://localhost:3000/nested
- http://localhost:3000/nested/
- http://localhost:3000/nested/index
- http://localhost:3000/nested/index.html

**`/foo.html` page**:

- http://localhost:3000/foo
- http://localhost:3000/foo/
- http://localhost:3000/foo.html
