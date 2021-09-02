# 📦 Truewallet Gift Vouchers with Golang

⠀⠀**Today**, people who make websites about purchasing goods or anything related to that transaction. They will have one of the transaction options, **Truewallet**, which is widely used in Thailand. However, we cannot use the core of the transaction in the normal way. **Unless you are a company and request to use the services of the official Truewallet for the core of the transaction.**


⠀⠀Until Truewallet has created a **Gift Voucher** to send money on any occasion through a link that Truewallet has generated. Just send the link to someone else and they will be able to receive your money.

**Of course, it's public. There is no encryption whatsoever, so developers can apply it to their websites.**

### ⚗️ functions/campaign.go (✅ 742 Bytes)
This is a file that sends data to Truewallet Gift Voucher service and it will callback back as Json.

```go
package functions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewRequestCampaign(mobile_number, campaign_code string) (string, error) {
	campaign_code := strings.Replace(campaign_link, "https://gift.truemoney.com/compaign?v=", "", -1)
	campaign_url := "https://gift.truemoney.com/campaign/vouchers/" + campaign_code + "/redeem"
	payload, _ := json.Marshal(map[string]string{"mobile": mobile_number})
	reader_buffer := bytes.NewBuffer(payload)
	response_campaign, err := http.Post(campaign_url, "application/json", reader_buffer)
	if err != nil {
		return "", err
	}
	defer response_campaign.Body.Close()

	body, _ := ioutil.ReadAll(response_campaign.Body)

    return string(body), nil
}
```

**🥴 The way I have it is to send all data in json format for easy and flexible management. You can select everything to be returned.**

```go
func NewRequestCampaign(mobile_number, campaign_code) { /* Code */ }
```

You can see that the **func NewRequestCampaign** It requires two parameters to send data to Truewallet:

- **mobile_number** Is the phone number of the registered with Truewallet and used for receiving money from campaign.
- **campaign_link** It is a link created to receive money from the account owner of that link. (Such as https://gift.truemoney.com/campaign?v={%#@})

------

### ⚗️ /example.go (✅ 1.54 Kilobytes)
Here is an example file I created to simulate as backend to pass data to front end via json.
```go
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
	Status struct {
		Code string `json:"code"`
	} `json:"status"`
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
```

In this code, I use a web framework called Fiber to manage the transmitted data. Convert to json format data and send data to frontend.

**You can execute this file for testing via:** `http://127.0.0.1:8008/transactions`
