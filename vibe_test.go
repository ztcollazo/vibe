package vibe_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/ztcollazo/vibe/example"
)

type VibeTestSuite struct {
	suite.Suite
	assert *assert.Assertions
	app    *fiber.App
}

func (s *VibeTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.app = example.RunApp()
}

func (s *VibeTestSuite) TestIndex() {
	testRoute(s.app, s.assert, "GET", "/", "Hello World!")
}

func (s *VibeTestSuite) TestNew() {
	testRoute(s.app, s.assert, "GET", "/new", "This is at GET /new")
}

func (s *VibeTestSuite) TestCreate() {
	testRoute(s.app, s.assert, "POST", "/", "This is at POST /")
}

func (s *VibeTestSuite) TestFind() {
	testRoute(s.app, s.assert, "GET", "/1", "This is at GET /:id with id of 1")
}

func (s *VibeTestSuite) TestEdit() {
	testRoute(s.app, s.assert, "GET", "/1/edit", "This is at GET /:id/edit with id of 1")
}

func (s *VibeTestSuite) TestUpdate() {
	testRoute(s.app, s.assert, "POST", "/1", "This is at POST /:id with id of 1")
}

func (s *VibeTestSuite) TestDestroy() {
	testRoute(s.app, s.assert, "DELETE", "/1", "This is at DELETE /:id with id of 1")
}

func (s *VibeTestSuite) TestCustom() {
	testRoute(s.app, s.assert, "GET", "/custom", "This is a custom route at GET /custom")
}

func TestVibe(t *testing.T) {
	suite.Run(t, new(VibeTestSuite))
}

func testRoute(app *fiber.App, assert *assert.Assertions, method, path, res string) {
	req := httptest.NewRequest(method, path, nil)
	resp, err := app.Test(req, 1)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)
	str, err := io.ReadAll(resp.Body)
	assert.Nil(err)
	assert.Equal(res, string(str))
	defer resp.Body.Close()
}
