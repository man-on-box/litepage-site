---
slug: initialize-new-litepage-instance
title: Create a new instance
section: Methods and usage
description: Learn about Litepage methods and their usage
---

# Litepage methods and usage

We've taken a look at a basic example, now lets go through each method and their use cases.

## Intializing a new instance

Create a new Litepage instance with `New` and passing your domain name, where the site will be accessed from.

```go
lp, err := litepage.New("hello-world.com")
```

The domain is used when building the sitemap, to make sure we can create fully qualified URLs.

### Initializing options

Optionally you can pass configuration options when creating the instance:

```go
lp, err := litepage.New("hello-world.com",
    litepage.WithDistDir("custom_dist"),
    litepage.WithBasePath("/custom-base"),
    litepage.WithPublicDir("custom_public"),
    litepage.WithoutSitemap(),
)
```

#### Options:

- `WithDistDir` - Specify a custom dist directory to be used when building the static site. Default value is `dist`.
- `WithBasePath` - Specify the base path of your site, if it is not the root of the domain (for example, if deploying to GitHub Pages). If set, all static assets and links should add the base as a prefix. The path should always start with a `/` and not end with a trailing slash (otherwise an error will be returned).
- `WithPublicDir` - Specify a custom public directory to be used to retrieve static assets when building or serving the static site. Default value is `public`.
- `WithoutSitemap` - Do not create a sitemap of your site. By default a sitemap.xml is created mapping all pages of the static site. Disable this if you do not want this, or if you want to create your own bespoke sitemap.

## Multiple instances

As you can configure each instance separately, you could create **multiple Litepage sites within the same project**. This can be useful if you are building multiple sites that share a lot of the same code and packages, like in a [monorepo](https://en.wikipedia.org/wiki/Monorepo) approach.

Of course this would require work in your continuous integration process, to make sure the sites are built and deployed to the correct places.
