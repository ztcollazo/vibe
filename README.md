# Vibe

**Always feel good about your Golang servers**.

## About

Vibe is a project that is meant to be used alongside [Fiber](https://gofiber.io), a fast and easy web server library for Go. It has a very similar API to Express. Vibe extends Fiber's functionality by adding controller functionality, designed to be Rails-like.

## API

See the [docs](https://go.dev/pkg/github.com/ztcollazo/vibe)

## Documentation

Here's an example of the API:

```go
package main

import (
  "log"

  "github.com/gofiber/fiber/v2"
  "github.com/ztcollazo/vibe"
)

type AppController struct {
  vibe.Controller[*AppController]
  Hello string
}

func (a *AppController) Setup(c *fiber.Ctx) {
  a.Hello = "Hello World!"
}

func (a *AppController) Index(c *fiber.Ctx) error {
  return c.SendString(a.Hello)
}

func main() {
  app := fiber.New()
  app.Route("/", vibe.CreateController(&AppController{}))
  log.Fatal(app.Listen(":3000"))
}
```

This is a basic example. You can view a (slightly) more in-depth example in the `example` folder As you can see, it uses an  `Index` method to route to `/`. The supported builtin routes are:

- `Index`: `GET /`
- `Show`: `GET /:id`
- `New`: `GET /new`
- `Create`: `POST /`
- `Edit`: `GET /:id/edit`
- `Update`: `POST /:id`
- `Destroy`: `DELETE /:id`

Defining custom routes is easy. You can create a `Routes` function on your controller, which returns a value of type `map[string]vibe.Route`. For example:

```go
func (a *AppController) Hello(c *fiber.Ctx) error {
  return c.SendString("This will be at GET `/hello`")
}

func (a *AppController) Routes() map[string]vibe.Route {
  return map[string]vibe.Route{
    "Hello": {
      Path: "/hello",
      Method: "GET",
    }
  }
}
```

It's that easy. You can even leave out a property such as `Path` or `Method` (but not both, so that the value is not the Zero value), and Vibe will guess it for you.

## Roadmap

- [X] Setup Middleware
- [X] Request Handlers
- [X] Default Paths
- [ ] Custom context and rendering

...and more that may come up in the future.

## License

Vibe is licensed under the MIT license. View the [LICENSE.txt](./LICENSE.txt) for more information.
