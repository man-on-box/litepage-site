---
slug: index
title: Installation
section: Getting started
description: Learn how to get started and add Litepage to your project
---

# Getting started with Litepage

Litepage is a lightweight library written in Go that simplifies building static sites. You declare the pages you want to create and it will build or serve your static site, along with any assets in your `/public` folder.

You can use this library in conjunction with any static site hosting provider like [GitHub Pages](https://pages.github.com/), [Cloudflare Pages](https://pages.cloudflare.com/) to get your static site up and running with minimal effort.

## Prerequisites

These are not hard requirements as we'll go over these later, but they are recommended.

- Basic knowledge of [Go](https://go.dev/)
- Basic knowledge of Go [HTML templating](https://pkg.go.dev/html/template)

## Installation

Add Litepage package to your project. There are no CLI's to use or config files setup.

```bash

go get github.com/man-on-box/litepage

```

## Project structure

Litepage has not much opinion on the folder layout of your project. The only folders Litepage interacts with by default are:

- `public/*` - to place your static assets (js, css, icons, images, etc.)
- `dist/*` - contains the outputted site when built, ready to be hosted

With these steps done, you are ready to start creating your static site!
