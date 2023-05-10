# Go clean architecture

This repository is an implementation of
the [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin) tutorial to learn Go
language.

The purpose of this project was to learn Go and to experiment:

* Clean architecture in Go.
* The [testify](https://github.com/stretchr/testify) package for the tests.
* [TDD](https://en.wikipedia.org/wiki/Test-driven_development) in Go.
* The [Gin Web Framework](https://github.com/gin-gonic/gin) (REST API).
* Dependency injection.

⚠️ Remark: as it was my first project in Go, Go best practices might be not respected.

## 🚀 Get started

To get started:

* Install the Go version defined in [.tool-versions](.tool-versions), I recommend to use [asdf](https://asdf-vm.com/).
* Install [make](https://www.gnu.org/software/make/#download).
* To test if the application works, run:

```shell
# In one Terminal
make start

# In another Terminal
make get-albums # => it should return a valid HTTP response
```

* Run `make help` to see all available commands.
