package main

import (
	"encoding/json"
	"log"
	"truewallet-gift-voucher-with-golang/functions"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Voucher struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

type Handler struct {
	Status Info `json:"status"`
}

type Info struct {
	Code string `json:"code"`
}

func main() {
	perform := fiber.New(fiber.Config{StrictRouting: true})
	perform.Use(logger.New())

	perform.Post("/transactions", func(c *fiber.Ctx) error {
		voucher := new(Voucher)
		if err := c.BodyParser(voucher); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"Code":    400,
				"Message": "Bad Request",
			})
		}

		redeemer, err := functions.NewRequestCampaign(voucher.Mobile, voucher.Code)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"Code":    400,
				"Message": "Bad Request",
			})
		}

		handler := new(Handler)
		if err := json.Unmarshal([]byte(redeemer), handler); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"Code":    422,
				"Message": "Unprocessable Entity",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Code":    handler.Status.Code,
			"Message": "<You can edit the message for your service.>",
		})
	})

	perform.Listen("127.0.0.1:8080")
}
