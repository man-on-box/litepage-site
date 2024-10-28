---
slug: build-or-preview-site
title: Build or preview
section: Methods and usage
description: Learn about Litepage methods and their usage
---

# Build or preview

We've seen the `Build` and `Serve` methods. Both methods explicitly build or serve your site respectively.

However, it is a common need to be able to serve your site locally while developing, and building it during CI without needing to change which method you call in your code.

For this you can take advantage of the `BuildOrServe` method, which will either (you guessed it) **build** or **serve** your site depending on environment variables.

```go
lp, _ := litepage.New("hello-world.com")

// ... add all your pages

err := lp.BuildOrServe()
```

By default, this will **build** your site as if you had called the `Build` method, however you can optionally `Serve` your site with the use of environment variables.

The following environment variables are checked when calling this method:

- `LP_MODE` - set this to `serve` to serve your site
- `LP_PORT` - set this to customise the port to serve your site on (default '3000')

## Example

Here is an example [Makefile](https://www.gnu.org/software/make/manual/make.html) where we are building or serving the site by specifying environment variables.

```make
build:
	# No env variables set here, so building by default
	go run cmd/main.go

serve:
	# LP_MODE is set to 'serve', so the site will be served on
	# default port 3000
	LP_MODE=serve go run cmd/main.go

dev:
	# LP_MODE is set to 'serve' and LP_PORT is set to 3001, so the
	# site will be served on port 3001. Useful if you're using a
	# hot reloading proxy like 'air'
	LP_MODE=serve LP_PORT=3001 air
```
