# helloserver

Very simple example using Go HTML templates for a server.
It has embedded HTML template and CSS file, so it
can run as a standalone binary without internet access.

To get dependencies and run, do

```sh
> go mod tidy
> go run .
```

To compile an executable:

```sh
> go build .
```

To run the tests, do

```sh
> go test ./...
```

You can then access the URLs:

* http://localhost:8888
* http://localhost:8888/healthz
* http://localhost:8888/debug/

Recommended literatures:

[Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/)
