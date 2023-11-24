# 201623-site

## How to install

Full instructions for installation are available in the USING.md file.
If you already know what you're doing, the instructions below should suffice.

Configuration for site is located in config.toml, and you'll need a
[MyRadio](https://github.com/UniversityRadioYork/MyRadio) API key with the
requisite permissions copied into a .myradio.key file in the same directory.

Then follow the Requirements below.

## Requirements

Requires [Go 1.6](https://golang.org/) to compile and run, along with `sassc` to
compile the SCSS files. You may use other SASS compilers if you wish, but be
prepared for unexpected results.

Alternatively, you can use Docker alone

## Running the site

### Without Docker

```bash
$ make run # Builds scss files, and runs the server
```

### With Docker :whale:

```bash
$ make build-docker-image #Builds the image, will only have to be re-run if you change the Dockerfile
$ make docker #Runs the image
```

## Editor Config

There is a handy editor config file included in this repo, most editors/IDE's have support for this either natively or through a plugin, see [here](http://editorconfig.org/#download).

## With Air (Live Refresh)

Air is a tool for live rebuilding on file edits, saving you having to manually kill and restart the running local server whenever you edit the Go or SCSS. An air config file for rebuilding the Go ans SCSS is included with this project

### Install Air

```bash
$ go install github.com/cosmtrek/air@latest
```

You may also have to manually set up the alias for it to run properly

```bash
$ alias air='$(go env GOPATH)/bin/air'
```

Once you've done this, you can just run:

```bash
$ air
```

and this should set up a live refreshing dev environment

### XDOTOOL for live browser refresh

Also included in this project is an XDOTOOL script for auto refreshing the browser when a change is made, this script is disabled by default as it only currently works on X11/Linux and requires a little setup.

First, install `XDOTOOL`, then, replace the line in .air.toml
` cmd = "make run"`
with
`cmd = "sh liverefresh.sh"`

The script is currently configured to work with Firefox, however, for other browsers, edit the line
`WID=$(xdotool search --name "Mozilla Firefox")`
to include your browser of choice
