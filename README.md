# HTTPSERVER

A helper CLI that create HTTP server from terminal.

## Installation

### Go

If you have Go 1.20+, you can directly install by running:

```
$ go install github.com/abdullahbagyapan/httpserver@latest
```

and the resulting binary will be placed at **$HOME/go/bin/httpserver**.

## Quick Start

```
$ httpserver -h
httpserver is a simple HTTP server CLI application, that written in Go.
This application is a tool to create basic HTTP server.

Usage:
  httpserver [flags]

Flags:
  -a, --addr ip            IP address for HTTP server (default 127.0.0.1)
  -d, --directory string   Directory for HTTP server (default "~/")
  -h, --help               help for httpserver
  -p, --port uint16        Port number for HTTP server (default 8080)
```

