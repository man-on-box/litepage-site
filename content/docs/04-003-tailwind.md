---
slug: tailwind
title: Tailwind CSS
section: Common recipes
description: Learn how to use Tailwind CSS with your Litepage project
---

# Using Tailwind CSS

[Tailwind CSS](https://tailwindcss.com/) is a hugely popular CSS framework for specifying styles within your HTML.

Note that even though Tailwind does provide a [standalone executable](https://tailwindcss.com/blog/standalone-cli), for ease of use locally and in your CI, this recipe will include installing Tailwind using **Node.js**.

> **Note:** You will need **npm** and **Node** on your machine before you get started.

## Installation

1. [npm init](https://docs.npmjs.com/cli/v6/commands/npm-init) your project (if you do not yet have a `package.json` file)
1. Run through the Tailwind [installation guide](https://tailwindcss.com/docs/installation)

Once setup, ensure you have configured your template paths to point to your HTML templates directory, and anywhere else you might specify class names that you want Tailwind to pickup (e.g. in your Markdown content).

## Usage

Now with Tailwind setup, you just need to ensure you run Tailwind **every time** you build your Go project. This ensures your CSS and outputted HTML stay in sync.

Take the below Makefile for example:

```make
build:
	@echo "Building for prod"
	@go run cmd/main.go
	@npx tailwind -i ./style/main.css -o ./dist/main.css --minify

serve:
	@npx tailwind -i ./style/main.css -o ./public/main.css & \
		LP_MODE=serve go run cmd/main.go

dev:
	@echo "Starting dev mode"
	@LP_MODE=serve LP_PORT=3001 air & \
		npx tailwind -i ./style/main.css -o ./public/main.css --watch
```

- `build`: We are building the site and CSS with Tailwind for production. We ensure the output destination is the dist directory and also minifying the outputted CSS.

- `serve`: We are creating the `main.css` file in the `/public` directory when we are building the site. This ensures the CSS is up to date, and the CSS is served when previewing the site (since Litepage hosts all files in `/public` when serving the site)

- `dev`: We start the app with Air, and run Tailwind with the `--watch` flag. This allows us to preview the site locally with hot reloading enabled, and also rebuilds the CSS files at the same time a change occurred

> **Note**: We would recommend to add `public/main.css` to your `.gitignore` to avoid adding your compiled CSS to your projects repository.
