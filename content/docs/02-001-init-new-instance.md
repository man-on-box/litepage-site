---
slug: initialize-new-litepage-instance
title: Create a new instance
section: Methods and usage
description: Learn about Litepage methods and their usage
---

# Litepage Methods and usage

We've taken a look at a basic example, now lets go through each method and their use cases.

## Intializing a new instance

To create a new Litepage instance:

```go

lp, err := litepage.New("hello-world.com")

```

Optionally you can pass configuration options when creating the instance:

```go
lp, err := litepage.New("hello-world.com",
    litepage.WithDistDir("custom_dist"),
    litepage.WithPublicDir("custom_public"),
    litepage.WithoutSitemap(),
)
```

**Options:**

- `WithDistDir` - Specify a custom dist directory to be used when building the static site. Default value is `dist`.
- `WithPublicDir` - Specify a custom public directory to be used to retrieve static assets when building or serving the static site. Default value is `public`.
- `WithoutSitemap` - Do not create a sitemap of your site. By default a sitemap.xml is created mapping all pages of the static site. Disable this if you do not want this, or if you want to create your own sitemap.
