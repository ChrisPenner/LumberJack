# LumberJack [![CircleCI](https://circleci.com/gh/ChrisPenner/LumberJack.svg?style=svg)](https://circleci.com/gh/ChrisPenner/LumberJack)
[Download Binary](https://github.com/ChrisPenner/LumberJack/releases/latest)

![screenshot](docs/screenshot.png)
![demo](docs/demo.gif)

## What is it?
A command-line interface log viewer

## Why did I build it
I got really sick of trying to find the things I was looking for in a big blob of unstructured text.
Printing my logs out to the terminal just wasn't sufficient anymore.

## What's it do?
Well, first and foremost it show your logs...

It can also:
- Stream logs into the viewer as they occur
- Filter to only show lines matching a given search term or regular expression
- Perform a search (from the latest logs upwards of course)
- Display logs from multiple servers side by side
- Display multiple views into the same log file
- Highlight text matching a given regular expression

## Install

The simplest method is to download a binary for your platform here: 
[Download Binary](https://github.com/ChrisPenner/LumberJack/releases/latest)

Or you can install from source:

Assuming `go` is installed:
```bash
$ go get github.com/chrispenner/lumberjack
$ $GOPATH/bin/lumberjack logs1 logs2
```



Put `$GOPATH/bin` on your `$PATH` to use the `lumberjack` command.

## Keybindings

### Log view:
- `<enter>`: Select a log-file for the current pane
- `<tab>`: Toggle the filters/highlighters side-pane
- `?` or `/`: Start a search
- `w`: Toggle text-wrapping
- `^h` and `^l`: Switch panes left and right respectively
- `<up>` and `<down>`: Scroll 1 line at a time
- `b` or `^u`: Scroll up half a screen
- `^d`: Scroll down half a screen
- `G`: Scroll to bottom (latest) logs
- `n`: Find next occurrance
- `N`: Find previous occurrance
- `1-4`: Display 1-4 panes respectively
- `<shift> + 0-9`: Toggle the respective filter/highlighter

### Filter/Highlighter pane:
- `<tab>`: Toggle the filters/highlighters side-pane
- `<enter>`: Edit the current modifier
- `<space>`: Toggle the current modifier
- `j`: Move down one modifier
- `k`: Move up one modifier
