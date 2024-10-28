---
slug: build-site
title: Build site
section: Methods and usage
description: Learn about Litepage methods and their usage
---

# Building your site

Once you have created all your pages, you can build your site with the `Build` method.

```go
lp, _ := litepage.New("hello-world.com")

// ... add all your pages

err := lp.Build()
```

This will build every page and place them in the `/dist` directory. As well, all your assets in the `/public` and a `sitemap.xml` (if set to build with the sitemap) will also be placed in the dist directory.

The final site in `/dist` directory can then be used with your preferred static site hosting service like [GitHub Pages](https://pages.github.com/), [CloudFlare Pages](https://pages.cloudflare.com/), etc. By following their documentation, it should be a trivial process to automate your build and host the outputted files during continuous integration.
