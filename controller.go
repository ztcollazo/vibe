package vibe

import (
	"reflect"

	"github.com/carlmjohnson/truthy"
	"github.com/gobeam/stringy"
	"github.com/gofiber/fiber/v2"
)

type CtrBase interface {
	Setup(ctx *fiber.Ctx)
	Routes() map[string]Route
}

type ictr[Ctr CtrBase] interface {
	CtrBase
	init(app fiber.Router, ctr Ctr)
}

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

var Defaults = map[string]Route{
	"Index": {
		Method: "GET",
		Path:   "/",
	},
	"New": {
		Method: "GET",
		Path:   "/new",
	},
	"Create": {
		Method: "POST",
		Path:   "/",
	},
	"Edit": {
		Method: "GET",
		Path:   "/:id/edit",
	},
	"Update": {
		Method: "POST",
		Path:   "/:id",
	},
	"Destroy": {
		Method: "DELETE",
		Path:   "/:id",
	},
	"Show": {
		Method: "GET",
		Path:   "/:id",
	},
}

type Controller[Ctr CtrBase] struct {
	ictr[Ctr]
	routemap map[string]Route
	ctr      Ctr
}

func (c Controller[Ctr]) Setup(ctx *fiber.Ctx) {}

func (c Controller[Ctr]) Routes() map[string]Route {
	return map[string]Route{}
}

func (c *Controller[Ctr]) createRoutes(routes map[string]Route) {
	c.routemap = make(map[string]Route)
	keys := []string{"Index", "New", "Show", "Create", "Edit", "Update", "Destroy"} // To keep the map in order
	v := reflect.ValueOf(c.ctr)
	for k, r := range routes {
		if m := v.MethodByName(k); m.IsValid() {
			t := r
			if !truthy.Value(t.Handler) {
				t.Handler = m.Interface().(func(*fiber.Ctx) error)
			}
			if !truthy.Value(t.Method) {
				t.Method = "GET"
			}
			if !truthy.Value(t.Path) {
				str := stringy.New(k)
				t.Path = "/" + str.KebabCase().ToLower()
			}
			c.routemap[k] = t
		}
	}

	for _, k := range keys {
		if m := v.MethodByName(k); m.IsValid() && !truthy.Value(c.routemap[k]) {
			r := Defaults[k]
			t := r
			t.Handler = m.Interface().(func(*fiber.Ctx) error)
			c.routemap[k] = t
		}
	}
}

func (c *Controller[Ctr]) init(app fiber.Router, ctr Ctr) {
	c.ctr = ctr
	c.createRoutes(ctr.Routes())
	app.Use(func(c *fiber.Ctx) error {
		ctr.Setup(c)
		return c.Next()
	})
	for _, v := range c.routemap {
		app.Add(v.Method, v.Path, v.Handler)
	}
}
