# MJML Server

Test implementation of mjml as a server

This project is just a test and is not production ready

The goal is to avoid shell exec in case you need to use mjml module but don't use nodejs as main tool

It also help in case you use docker (https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices/)

## Server

```
hostname: localhost  
port: 8686
```

Run server (folder server)

`$` `node server.js`

### Protocol

```
[client] send mjml -> | (uint32 - Big Endian) size of the mjml template | (bytes) template mjml
[client] receive html -> | (uint32 - Big Endian) size of the html template | (bytes) template html
```

## Client

The clients read the template from "template.mjml" and request the server to convert it, then write the result in "output.html"

### Go

Run go client (folder client-go):

`$` `go run main.go`

### Ruby

Run ruby client (folder client-ruby):

`$` `ruby main.rb`
