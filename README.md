## Overview
`quack` is a tiny Go web server for testing endpoints inside networks that don't allow [RequestBin](https://requestb.in/). It listens on a designated port and responds to HTTP `GET`, `POST`, `PUT`, and/or `DELETE` requests with a witty, duck-themed message while logging the requests and their parameters/payloads to the console.

## Installation
```
go get github.com/jessemillar/quack
```

## Usage
```
usage: quack [<flags>]

Flags:
      --help         Show context-sensitive help (also try --help-long and --help-man).
  -p, --port="8000"  The port to listen on.
```

Use [Postman](https://www.getpostman.com/) (or similar) to send an HTTP `GET`, `POST`, `PUT`, or `DELETE` request to the server.
