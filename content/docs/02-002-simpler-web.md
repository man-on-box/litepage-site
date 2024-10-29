---
slug: a-simpler-web
title: A simpler web
section: Core concepts
description: Learn how Litepage aims to make building static sites simpler
---

# Creating a simpler web

Frameworks come and go. They evolve, they grow, they compete and they require varying degrees of knowledge to use and maintenance to keep up to date.

There are a plethora of frameworks that can build static sites, my favourite being [Astro](https://astro.build/), along with [Next.js](https://nextjs.org/), [Gatsby](https://www.gatsbyjs.com/), [Jekyll](https://jekyllrb.com/docs/), [Eleventy](https://www.11ty.dev/), [Hugo](https://gohugo.io/)... and this is just the [tip of the iceberg](https://jamstack.org/generators/).

## The cost of frameworks

Choosing a framework for your project is a bit like getting into a relationship. It requires attention, learning and long-term commitment. Depending on the volatility of said framework, you may incur technical debt over time, especially if you don't attend to it regularly.

If I could sketch this scenario in a chart, it would look something like:

![Image of productivity vs tech debt over time](/img/with-frameworks.webp)

By comparison, languages evolve and change at a much _slower_ rate, and should be a solid foundation to build on top of. I love [Go](https://go.dev/) for its simplicity and comprehensive standard library with built in HTML templating and web server.

Why, then, when creating a simple static site, must I reach for frameworks that introduce domain-specific languages, ever-changing APIs, new features, and potentially hundreds or thousands of dependencies into my seemingly innocent project?

Can I not just use the tools the language provides, and opt to bring in libraries when I need them? This is when Litepage was conceived.

## Isn't Litepage just another framework?

Well yes, and no. I wanted a library to handle the most common tasks when building a static site, mainly:

- Create the static files
- Copy static assets
- Be able to preview the site

And that is all. The rest, I wanted to be able to pick and choose on a per project basis. Do I want to use [Tailwind](https://tailwindcss.com/)? Sure. Do I just want to use JSON files as my source of truth? Sure. Do I want to parse Markdown files? Sure. Do I want to create a blog? No worries.

To keep things simple, Litepage works with just the standard `io.Writer` interface, meaning you are in control on how you parse your data and templates. If you wanted to remove Litepage from your project at a later date, you could just swap it out and write the code manually to create your static site.

## Isn't this extra effort?

Yes! It's extra effort to get started. There is no _magic_ here, just simplicity. While it is extra effort to get started, later down the line you should incur less techincal debt, as it's very unlikely Go, HTML or CSS will have changed much.

For small projects that do not require much complexity, keeping things simple will pay dividends as time goes by.

## Sounds like NIH

[Not invented here](https://en.wikipedia.org/wiki/Not_invented_here) syndrome is a real thing. I'm an advocate to build things when it makes sense to do so, either because **it's something simple**, or because **you need something bespoke**. For Litepage, it fell under the former.
