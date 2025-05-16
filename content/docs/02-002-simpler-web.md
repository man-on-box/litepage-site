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

To go back to the basics. Litepage is more of a philosophy than a library. It is an approach to build your sites in native Go, leveraging its standard library. Litepage can then build your site for production, and serve it whilst you are developing.

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

## What about Javascript?

The problem is Javascript is that to do anything complex, you have to reach for packages. These packages in turn rely on other packages, and those on others, etc, etc. This is why modern JS frameworks rely on _hundreds_ of dependencies to function.

Javascript is a cool language, but its even better when its used as it was intended. Forget Typescript, transpilation and build steps. Just write your Javascript as it was intended, to enrich funcionality to your HTML elements in your browser.

## What about not so simple sites?

More complex sites deserve more complex stacks. Are you building the next Netflix, Gmail, Shopify or Jira? Then you need a more complex stack, and this is what stacks like [Astro](https://astro.build/), [NextJs](https://nextjs.org/), [Laravel](https://www.jetbrains.com/phpstorm/laravel/) and others are made for.
