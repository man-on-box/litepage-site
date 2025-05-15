---
slug: a-simpler-web
title: A simpler web
section: Core concepts
description: Learn how Litepage aims to make building static sites simpler
---

# Creating a simpler web

If your site is serving static content with little client side interaction, you **don't need** a server and you **don't need** a framework.

The majority of websites out there are simple sites serving static content, and too many of them are hosted on servers, wasting computing power and incurring unnecessary costs.

As well, it has become the norm to use frameworks that bring in _hundreds_ of dependencies, just to build and ship HTML over the web.

## So what's the alternative?

To go back to the basics. Litepage is more of a philosophy than a library. It is an approach to build your sites in native Go, leveraging its standard library. Litepage can then build your sites for production, and serve them whilst you are developing.

## What about performance?

Frameworks promise performance, but the best way to build performant sites is to not ship bloat. A site does not inherently run _faster_ on a users machine because you used a certain framework, but it can run _slower_.

The best way to give your users peformance, is to **ship the bare minimum HTML, CSS and JS required**.

## Your actual project life cycle.

The reality is that **most of your time is spent on maintenance, not development.**

Frameworks promise immediate productivity, and in most cases this is true. But once your site is built and done, for most of its life it will sit there accruing technical debt and security vulnerabilities.

With this in mind, you can build your site for the long term. This means **spending more time during development**, but **less time overall**. As the weeks, months and years go by, your code stays as relevant as it was when you first wrote it.

## Why Go?

So why use Golang? It's the obvious choice to build something stable and simple. With its built-in HTML templating and web server, you can go far without needing _any_ external dependencies.

As well, Golang is a really useful language to learn, especially if you are a frontend developer. It can be tricky at first, but once you get the hang of it you can quickly start writing performant and stable web projects which require very few dependencies and little maintenance.

## --- To Delete / Old content ---

Frameworks come and go. They evolve, they grow, they compete and they require varying degrees of knowledge to use and maintenance to keep up to date.

There are a plethora of frameworks that can build static sites, my favourite being [Astro](https://astro.build/), along with [Next.js](https://nextjs.org/), [Gatsby](https://www.gatsbyjs.com/), [Jekyll](https://jekyllrb.com/docs/), [Eleventy](https://www.11ty.dev/), [Hugo](https://gohugo.io/)... and this is just the [tip of the iceberg](https://jamstack.org/generators/).

## The cost of frameworks

Choosing a framework for your project is a bit like getting into a relationship. It requires attention, learning and long-term commitment. Depending on the volatility of said framework, you may incur technical debt over time, especially if you don't attend to it regularly.

If I could sketch this scenario in a chart, it would look something like:

<img class="max-h-96 mx-auto" src="/img/prod-chart-light.svg" alt="Image of productivity vs tech debt over time">

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
