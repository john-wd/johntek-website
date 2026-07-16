# Johntek Consulting website

This is the source for the Johntek Consulting website. It is a Hugo site with custom layouts, content sections, blog posts, resume data, static images and a Go command for site management and hugo wrapper.

Most of the day to day work happens in `content`, `layouts`, `assets`, `data` and `static`. The generated site lands in `public` after a build, as it is expected from Hugo.

## What is in here

1. `content` has the pages, blog posts and landing page sections.
2. `layouts` has the Hugo templates that turn the content into pages.
3. `assets` has CSS, scripts, icons and source images that Hugo can process. CSS is autodetected and merged during build time.
4. `static` has files copied directly into the published site.
5. `data` has structured content such as the resume and testimonials.
6. `cmd` has the Go command entry point for the local site tooling.

The biggest difference with vanilla Hugo is that the landingpage layout ([layouts/_default/landingpage.html](layouts/_default/landingpage.html)) is modular and lets you compose sections from reusable components.

See [content/landing-sections](content/landing-sections) for an example how the default index.html is declared there. The ordering of sections is done by the `weight` field in the front matter.

## First setup

If just using Hugo, install it here https://gohugo.io/installation/linux/. It is available in most package managers, so use brew, apt or pacman to install it. Then run the regular `hugo` commands you see below.

```sh
go mod download
```

If you want to install the command utilities as well, first install Go, then download the Go dependencies. After this you can run my custom commands by typing `go run . <command>`.

## Serve it locally

The most direct way to work on the site is Hugo's local server:

```sh
hugo serve
```

If you want to run Hugo through the project wrapper instead, use:

```sh
go run . hugo serve
```

Either way, Hugo will print the local URL in the terminal. It is usually `http://localhost:1313/`.

## Build it

For a normal local build:

```sh
hugo
```

For the production style build used by the deploy workflow:

```sh
hugo --minify
```

The output goes into `public`.

## About the commands

The command code lives in `cmd/root.go`. It defines a `webctl` using Cobra and adds a `hugo` subcommand. That subcommand passes everything after `hugo` straight through to Hugo, so this:

```sh
go run . hugo server
```

behaves like this:

```sh
hugo server
```

Right now the wrapper is small, but it gives the repo a place to grow site related commands without scattering scripts around the project.

## Deploys

The GitHub workflow in `.github/workflows/deploy.yaml` builds with:

```sh
hugo --minify
```

Then it syncs `public` to Hostinger over SSH. Deploys run on pushes to `main`, on a schedule and by manual dispatch.
