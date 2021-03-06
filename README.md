## Overview

[![Man Hours](https://img.shields.io/endpoint?url=https%3A%2F%2Fmh.jessemillar.com%2Fhours%3Frepo%3Dhttps%3A%2F%2Fgithub.com%2Fjessemillar%2Fquack.git)](https://jessemillar.com/r/man-hours)

`quack` is a tiny Go web server for testing endpoints inside networks that don't allow [RequestBin](https://requestb.in/). It listens on a designated port and responds to HTTP `GET`, `POST`, `PUT`, and/or `DELETE` requests with a witty, duck-themed message while logging the requests and their parameters/payloads to the console.

## Installation
```
go get -v -u github.com/jessemillar/quack
```

## Usage
```
usage: quack [<flags>]

Flags:
      --help         Show context-sensitive help (also try --help-long and --help-man).
  -p, --port="8000"  The port to listen on.
```

Use [Postman](https://www.getpostman.com/) (or similar) to send an HTTP `GET`, `POST`, `PUT`, or `DELETE` request to the server.
