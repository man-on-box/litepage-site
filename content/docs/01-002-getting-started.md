---
slug: installation
title: Installation
section: Getting started
description: Learn how to get started and add Litepage to your project
---

# Adding Litepage to your project

## Prerequisites

These are not hard requirements as you can learn as you go, but they are recommended.

- Basic knowledge of [Go](https://go.dev/)
- Basic knowledge of Go [HTML templating](https://pkg.go.dev/html/template)
- Knowledge of HTML, CSS, JS (no ready-made templates here)

## Bootstrap a new project

Right now we are **building a CLI** that can bootstrap a new project for you, with some basic optional configuration such as with Markdown or Tailwind CSS support. This isn't ready yet, so for now you can add the library manually to a new or existing project.

## Add to new/existing project

Add Litepage package to your project as you would any other package.

```bash
go get github.com/man-on-box/litepage
```

## Project structure

Litepage has not much opinion on the folder layout of your project. The only folders Litepage interacts with by default are:

- `public/*` - to place your static assets (js, css, icons, images, etc.)
- `dist/*` - contains the outputted site when built, ready to be hosted

### Example project structure

For a small project, it may look something like this, but you are free to structure however you wish.

```
my-cool-site
└── content
  ├── some-json.json
  └── some-markdown.md
└── dist
  └── *built site files*
└── public
  ├── favicon.ico
  ├── logo.svg
  └── styles.css
└── view
  └── homepage.html
└── main.go
└── go.mod
└── go.sum
```

With these steps done, you are ready to start creating your static site!
