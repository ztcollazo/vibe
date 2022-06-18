package vibe

import "github.com/gofiber/fiber/v2"

func CreateController[Ctr ictr[Ctr]](ctr Ctr) func(fiber.Router) {
	return func(r fiber.Router) {
		ctr.init(r, ctr)
	}
}
