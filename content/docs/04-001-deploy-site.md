---
slug: deploy-site
title: Deploy site
section: Common recipes
description: Common recipes to get started and solving common usecases
---

# Deploying your site

Once you have built your site, you need to deploy it somewhere. It should be possible to deploy to any static site provider that supports a Go runtime, and we can provide some guides here.

## Cloudflare Pages

[Cloudflare Pages](https://developers.cloudflare.com/pages/get-started/git-integration/) offers an easy way to deploy your static sites, with preview deployments and basic analytics out of the box.

1. Push your project to [Github](https://github.com/) or [GitLab](https://gitlab.com/)
1. Login to the [Cloudflare dashboard](https://dash.cloudflare.com/) and select your account
1. Select **Workers & Pages** > **Create** > **Pages** > **Connect to Git**

Here you will be prompted to sign in with your Git provider. Once connected, select the repository that contains your static siteproject and click **Begin Setup**

Once connected, it is as simple as selecting your build command to build your project and site (e.g. `make build` or `go run main.go`), and the build output directory, which by default is `/dist`.

From here you can deploy your site everytime you commit to your main branch, as well as get automatic preview URLs when pushing up new branches and pull requests. You can [read their docs](https://developers.cloudflare.com/pages/) for more information.

## GitHub Pages

GitHub Pages allows you to host your static site straight from GitHub. This is convenient especially if you are already using GitHub to put your project. It leverages [GitHub Actions](https://github.com/features/actions) to build and deploy your site automatically.

1. Create a repository named _username_.github.io
