---
slug: creating-a-new-site
title: Create a new site
section: Getting started
description: Learn how to get started and add Litepage to your project
---

# Creating a new site with Litepage

Lets start first with a basic example, then we'll walk through and explain each method below.

**Basic example:**

```go
func main() {
	lp, _ := litepage.New("hello-world.com")

	lp.Page("/index.html", func(w io.Writer) {
		html := "<h1>Hello, World!</h1>"
		t := template.Must(template.New("helloWorld").Parse(html))
		t.Execute(w, nil)
	})

	lp.Build()
}
```

- We create a new litepage instance with `litepage.New` for our pretend site `hello-world.com`.
- We add a new `/index.html` page with `lp.Page`. In this page, we are creating a new html template of `<h1>Hello, World!</h1>` and writing it to the index page.
- Lastly we build the site with `lp.Build`, which builds our site into the `/dist` directory. This directory will contain the `index.html` file, as well as any assets you have in your `/public` directory.

**And that's it!**

It really is that simple. Our site is now ready to deploy on a static site provider somewhere.

We could leave it here, but most likely you will want to build something slightly more ambitious and learn more about the development flow. For this we can learn more about the methods individually in the next section.
